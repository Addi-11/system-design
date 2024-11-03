package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Agent struct {
	ID string
	Name string
	OrderID sql.NullString
	IsReserved bool
}

func reserveAgent() (*Agent, error) {
	txn, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	// select first available delivery agent
	row := txn.QueryRow(`
		SELECT id, name, order_id, is_reserved FROM agents
		WHERE is_reserved = false AND order_id IS NULL
		LIMIT 1
		FOR UPDATE
	`)

	if row.Err() != nil {
		txn.Rollback()
		return nil, row.Err()
	}

	var agent Agent
	err = row.Scan(&agent.ID, &agent.Name, &agent.OrderID, &agent.IsReserved)
	if err != nil {
		txn.Rollback()
		return nil, errors.New("no delivery agent available")
	}

	// reserve the agent
	_, err = txn.Exec(`
		UPDATE agents SET is_reserved = true
		WHERE id = ?
	`, agent.ID)
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

func reserveHandler(c *gin.Context) {
	agent, err := reserveAgent()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ReserveAgentResponse{AgentID: agent.ID})
}

type ReserveAgentResponse struct {
	AgentID string `json:"agent_id"`
}
