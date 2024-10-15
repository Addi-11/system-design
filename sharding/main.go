package main

import (
	"fmt"
	"os"
	"database/sql"
	"net/http"
	"log"
	"strconv"
	
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DBs []*sql.DB

// create 2 connections to our DB simulating multiple shards
func init(){
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "test_db"

	// connection string to the DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	_db1, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = _db1.Ping()
	if err != nil {
		log.Fatal("Error pinging the database 1: ", err)
	}


	DBs = append(DBs, _db1)

	_db2, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = _db2.Ping()
	if err != nil {
		log.Fatal("Error pinging the database 1: ", err)
	}


	DBs = append(DBs, _db2)
}

// given a user, map it to the shard it belongs to
// implement routing logic: static mapping
// simple hashmap here: even users DB 0, odd users DB 1
func getShardIdx(userID string) int{
	id, err := strconv.Atoi(userID)
	if err != nil{
		fmt.Println("Invalid userID", userID)
		return -1
	}

	index := id % 2

	fmt.Printf("UserID: %s using DB: %d\n", userID, index)
	return index
}


func getUser(c *gin.Context){
	userID := c.Param("userID")

	// for respective DB for the user
	DB := DBs[getShardIdx(userID)]

	// query the DB to get user-details
	row := DB.QueryRow("SELECT name FROM users WHERE id = ?", userID)

	// scan the query result to get the varchar name
	var name string
	err := row.Scan(&name)

	if err != nil{
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// return  health status
	c.JSON(http.StatusOK, gin.H{"userID": userID, "name": name})
}

func main() {
	r := gin.Default()

	// route to get user by USER_ID
	r.GET("/user/:userID", getUser)

	if err := r.Run(":8080"); err != nil{
		panic(err)
	} 

}
