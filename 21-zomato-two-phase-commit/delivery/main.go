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


func init(){
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "zomato"

	// connection string to the DB
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
}

type Agent struct{
	ID string
	Name string
	OrderID sql.NullString
	IsReserved bool
}

func reserveAgent() (*Agent, error){
	txn, _ := DB.Begin()

	// select the first available delivery agent
	row := txn.QueryRow(`
		SELECT id, name, order_id, is_reserved from agents
		WHERE is_reserved is false and order_id is NULL
		LIMIT 1
		FOR UPDATE
	`)
	if row.Err() != nil{
		txn.Rollback()
		return nil, row.Err()
	}

	var agent Agent
	err := row.Scan(&agent.ID, &agent.Name, &agent.OrderID, &agent.IsReserved)

	if err != nil{
		txn.Rollback()
		return nil, errors.New("no delivery agent available")
	}

	// reserve the agent
	_, err = txn.Exec(`
		UPDATE agents SET is_reserved = true
		WHERE id = ?
	`, agent.ID)
	if err != nil{
		txn.Rollback()
		return nil, err
	}
	err = txn.Commit()
	if err != nil{
		return nil, err
	}
	return &agent, nil
}

func bookAgent(orderID string) (*Agent, error) {
	txn, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	// Select the first reserved agent
	row := txn.QueryRow(`
		SELECT id, name, order_id, is_reserved FROM agents
		WHERE is_reserved = true AND order_id IS NULL
		LIMIT 1
		FOR UPDATE
	`)

	var agent Agent
	err = row.Scan(&agent.ID, &agent.Name, &agent.OrderID, &agent.IsReserved)
	if err == sql.ErrNoRows {
		txn.Rollback()
		return nil, errors.New("no delivery agent available")
	} else if err != nil {
		txn.Rollback()
		return nil, err
	}

	_, err = txn.Exec(`
		UPDATE agents SET is_reserved = false, order_id = ?
		WHERE id = ?
	`, orderID, agent.ID)
	if err != nil {
		txn.Rollback()
		return nil, err
	}

	// Commit the transaction
	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

type ReserveAgentResponse struct {
	AgentID string `json:"agent_id"`
}

type BookAgentRequest struct {
	OrderID string `json:"order_id" binding:"required"`
}

type BookAgentResponse struct {
	AgentID string `json:"agent_id"`
	AgentName string `json:"agent_name"`
	OrderID string `json:"order_id"`
}

func main(){
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil{
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

	// reserve agent route
	r.POST("/delivery/agent/reserve", func(c *gin.Context){
		agent, err := reserveAgent()
		if err != nil{
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, ReserveAgentResponse{AgentID: agent.ID})
	})

	// book agent route
	r.POST("/delivery/agent/book", func(c *gin.Context){
		var req BookAgentRequest
		if err := c.ShouldBindJSON(&req); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		agent, err := bookAgent(req.OrderID)
		if err != nil{
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, BookAgentResponse{AgentID: agent.ID, AgentName: agent.Name, OrderID: req.OrderID})
	})

	// Start server
	fmt.Println("Delivery agent server running on port 8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}