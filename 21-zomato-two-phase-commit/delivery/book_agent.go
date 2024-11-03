package main

import (
	"database/sql"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
)

func bookAgent(orderID string) (*Agent, error) {
	txn, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	// select the first reserved agent
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

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	return &agent, nil
}

type BookAgentRequest struct {
	OrderID string `json:"order_id" binding:"required"`
}

type BookAgentResponse struct {
	AgentID   string `json:"agent_id"`
	AgentName string `json:"agent_name"`
	OrderID   string `json:"order_id"`
}

func bookHandler(c *gin.Context) {
	var req BookAgentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	agent, err := bookAgent(req.OrderID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, BookAgentResponse{AgentID: agent.ID, AgentName: agent.Name, OrderID: req.OrderID})
}
