package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Packet struct {
	ID         string
	FoodID     int
	OrderID    sql.NullString
	IsReserved bool
	PacketName string
}

func reserveFood(foodID int) (*Packet, error) {
	txn, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	// select first available food packet
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

	// reserve the food packet
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

func reserveFoodHandler(c *gin.Context) {
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
}

type ReserveFoodRequest struct {
	FoodID int `json:"food_id" binding:"required"`
}

type ReserveFoodResponse struct {
	PacketID string `json:"packet_id"`
}
