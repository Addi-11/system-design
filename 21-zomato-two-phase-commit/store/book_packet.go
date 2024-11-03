package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func bookFood(orderID string, foodID int) (*Packet, error) {
	txn, err := DB.Begin()
	if err != nil {
		return nil, err
	}
	// get food packet name
	var packetName string
	row := txn.QueryRow(`SELECT name from foods WHERE id = ?`, foodID)
	err = row.Scan(&packetName)
	if err != nil {
		return nil, errors.New("food item not found with id")
	}

	// select first reserved food packet
	row = txn.QueryRow(`
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
	foodPacket.PacketName = packetName

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

func bookFoodHandler(c *gin.Context) {
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
	c.JSON(http.StatusOK, BookFoodResponse{PacketID: foodPacket.ID, OrderID: req.OrderID, PacketName: foodPacket.PacketName})
}

type BookFoodRequest struct {
	OrderID string `json:"order_id" binding:"required"`
	FoodID  int    `json:"food_id" binding:"required"`
}

type BookFoodResponse struct {
	PacketID   string `json:"packet_id"`
	OrderID    string `json:"order_id"`
	PacketName string `json:"packet_name"`
}
