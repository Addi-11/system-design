package main

import (
	"fmt"
	"sync"
	"time"
	"database/sql"
	"os"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser = os.Getenv("DBUSER")
	dbPass = os.Getenv("DBPASS")
	dbHost = "127.0.0.1:3300"
	dbName = "test_db"

	// connection string to the DB
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
)

// open connection to the DB
func createNewConnection() (*sql.DB, error){
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	return db, nil
}

// simulate database operation
func simulateDBOperation (db *sql.DB) error {
	_, err := db.Exec("SELECT SLEEP(0.01);")
	return err
}

func benchmarkNonPool(n int) (time.Duration, error){
	start := time.Now()
	var wg sync.WaitGroup
	var once sync.Once
	var operr error

	helper := func (){
		defer wg.Done()
		conn, err := createNewConnection()
		if err != nil {
			once.Do(func() {operr = err})
			return
		}
		defer conn.Close()
		if err := simulateDBOperation(conn); err != nil {
			once.Do(func(){
				operr = err
			})
		}
	}

	for i:= 0; i < n; i++ {
		wg.Add(1)
		go helper()
	}
	// wait for all goroutines to complete
	wg.Wait()
	return time.Since(start), operr
}

func main(){
	tests := []int{10, 100, 200, 300, 500, 1000}
	
	for _, n := range tests {
		elapsed, err := benchmarkNonPool(n)

		if err != nil {
			fmt.Printf("Error running non-pool benchmark for %d threads: %v\n", n, err)
		} else {
			fmt.Printf("Non-pool time for %d threads: %v\n", n, elapsed)
		}
	}
	
}