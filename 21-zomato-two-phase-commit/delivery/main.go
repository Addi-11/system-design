package main

import (
	"database/sql"
	"fmt"
	"log"
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

func main() {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("error pinging the DB.")
	}
	defer DB.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/delivery/agent/reserve", reserveHandler)

	r.POST("/delivery/agent/book", bookHandler)

	// start server
	fmt.Println("Delivery agent server running on port 8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}