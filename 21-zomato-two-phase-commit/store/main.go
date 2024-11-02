package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var dsn string

func init() {
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "zomato"

	// Connection string to the DB
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
}

type Packet struct {
	ID         string
	FoodID     int
	OrderID    sql.NullString
	IsReserved bool
}

func reserveFood(foodID int) (*Packet, error) {
	txn, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	// Select the first available food packet
	row := txn.QueryRow(`
		SELECT id, food_id, order_id, is_reserved FROM packets
		WHERE food_id = ? AND is_reserved = false AND order_id IS NULL
		LIMIT 1
		FOR UPDATE
	`, foodID)

	var foodPacket Packet
	err = row.Scan(&foodPacket.ID, &foodPacket.FoodID, &foodPacket.OrderID, &foodPacket.IsReserved)
	if err == sql.ErrNoRows {
		txn.Rollback()
		return nil, errors.New("no food packet available")
	} else if err != nil {
		txn.Rollback()
		return nil, err
	}

	// Reserve the food packet
	_, err = txn.Exec(`
		UPDATE packets SET is_reserved = true
		WHERE id = ?
	`, foodPacket.ID)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}
	return &foodPacket, nil
}

func bookFood(orderID string, foodID int) (*Packet, error) {
	txn, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	// Select the first reserved food packet
	row := txn.QueryRow(`
		SELECT id, food_id, order_id, is_reserved FROM packets
		WHERE food_id = ? AND is_reserved = true AND order_id IS NULL
		LIMIT 1
		FOR UPDATE
	`, foodID)

	var foodPacket Packet
	err = row.Scan(&foodPacket.ID, &foodPacket.FoodID, &foodPacket.OrderID, &foodPacket.IsReserved)
	if err == sql.ErrNoRows {
		txn.Rollback()
		return nil, errors.New("no reserved food packet available")
	} else if err != nil {
		txn.Rollback()
		return nil, err
	}

	_, err = txn.Exec(`
		UPDATE packets SET is_reserved = false, order_id = ?
		WHERE id = ?
	`, orderID, foodPacket.ID)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}
	return &foodPacket, nil
}

type ReserveFoodRequest struct {
    FoodID int `json:"food_id" binding:"required"`
}

type ReserveFoodResponse struct {
	PacketID string `json:"packet_id"`
}

type BookFoodRequest struct {
    OrderID string `json:"order_id" binding:"required"`
    FoodID  int    `json:"food_id" binding:"required"`
}

type BookFoodResponse struct {
	PacketID string `json:"packet_id"`
	OrderID  string `json:"order_id"`
}

func main() {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil{
		log.Fatal("error pinging the DB.")
	}
	defer DB.Close()

	// Setup gin router
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Reserve food route
	r.POST("/food/reserve", func(c *gin.Context) {
		var req ReserveFoodRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		foodPacket, err := reserveFood(req.FoodID)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, ReserveFoodResponse{PacketID: foodPacket.ID})
	})

	// Book food route
	r.POST("/food/book", func(c *gin.Context) {
		var req BookFoodRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}
	
		foodPacket, err := bookFood(req.OrderID, req.FoodID)
		if err != nil {
			log.Printf("Error booking food for order %s: %v", req.OrderID, err)
			c.JSON(http.StatusConflict, gin.H{"error": "order not placed: could not assign food to the order"})
			return
		}
		c.JSON(http.StatusOK, BookFoodResponse{PacketID: foodPacket.ID, OrderID: req.OrderID})
	})

	// Start server
	fmt.Println("Food reservation server running on port 8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
