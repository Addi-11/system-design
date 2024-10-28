package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "test_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Check database connection
	if err := DB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}
}

// mimic EC2 instance creation
func createEC2(serverID int) {
	fmt.Println("creating server")
	_, err := DB.Exec("UPDATE p_servers SET status = 'TODO' WHERE server_id = ?;", serverID)
	if err != nil {
		log.Printf("Error updating server status to TODO: %v", err)
		return
	}

	time.Sleep(5 * time.Second)
	_, err = DB.Exec("UPDATE p_servers SET status = 'IN PROGRESS' WHERE server_id = ?;", serverID)
	if err != nil {
		log.Printf("Error updating server status to IN PROGRESS: %v", err)
		return
	}
	fmt.Println("server creation in progress")

	time.Sleep(5 * time.Second)
	_, err = DB.Exec("UPDATE p_servers SET status = 'DONE' WHERE server_id = ?;", serverID)
	if err != nil {
		log.Printf("Error updating server status to DONE: %v", err)
		return
	}
	fmt.Println("server creation completed")
}

func main() {
	ge := gin.Default()

	ge.POST("/servers", func(ctx *gin.Context) {
		// Here, you might want to call createEC2 asynchronously if it's intended to run in the background
		go createEC2(1) // Replace `1` with actual serverID as needed
		ctx.JSON(200, map[string]interface{}{"submitted": "ok"})
	})

	// short polling: The client sends a request at regular intervals to check for updates.
	ge.GET("/short/status/:server_id", func(ctx *gin.Context) {
		serverID := ctx.Param("server_id")

		var status string
		row := DB.QueryRow("SELECT status FROM p_servers WHERE server_id = ?;", serverID)
		err := row.Scan(&status)
		if err != nil {
			log.Printf("Error querying status: %v", err)
			ctx.JSON(500, map[string]interface{}{"error": "Internal server error"})
			return
		}

		ctx.JSON(200, map[string]interface{}{"status": status})
	})

	// long polling: the client sends single request, and the server holds the connection open until there’s an update or status change.
	ge.GET("/long/status/:server_id", func(ctx *gin.Context) {
		serverID := ctx.Param("server_id")
		currentStatus := ctx.Query("status")

		var status string
		for {
			row := DB.QueryRow("SELECT status FROM p_servers WHERE server_id = ?;", serverID)
			err := row.Scan(&status)
			if err != nil {
				log.Printf("Error querying status: %v", err)
				ctx.JSON(500, map[string]interface{}{"error": "Internal server error"})
				return
			}

			if currentStatus != status {
				break
			}
			time.Sleep(1 * time.Second)
		}

		ctx.JSON(200, map[string]interface{}{"status": status})
	})

	ge.Run(":9000")
}
