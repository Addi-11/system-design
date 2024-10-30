package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"log"

	"github.com/google/uuid"
	_ "github.com/go-sql-driver/mysql"
)

var dsn string

func init() {
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "test_db"

	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
}

func int_table(db *sql.DB, n int){
	for i :=0 ; i<n; i++{
		_, err := db.Exec("INSERT INTO table_int (age) VALUES (?)", i%100)
		if err != nil{
			panic(err)
		}
	}
}

func uuid_table(db *sql.DB, n int){
	for i:=0; i<n; i++{
		_, err:= db.Exec("INSERT INTO table_uuid (id, age) VALUES (?, ?)", uuid.New().String(), i%100)
		if err != nil{
			panic(err)
		}
	}
}

func createIndex(db *sql.DB, tableName, indexName, columnName string) {
	start := time.Now()
	query := fmt.Sprintf("CREATE INDEX %s ON %s(%s);", indexName, tableName, columnName)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create index on %s: %v", tableName, err)
	}
	fmt.Printf("Index creation on %s (%s): %v\n", tableName, columnName, time.Since(start))
}

func main(){
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	N := []int{1, 10, 100, 1000, 5000, 10000, 50000, 1000000}

	for _, n := range(N){
		// Benchmarking for INT primary key
		start := time.Now()
		int_table(db, n)
		fmt.Printf("Time taken for INT table n=%d: %v\n", n, time.Since(start))

		// Benchmarking for UUID primary key
		start = time.Now()
		uuid_table(db, n)
		fmt.Printf("Time taken for UUID table n=%d: %v\n", n, time.Since(start))
	}

	// Measure time to create index on table_int
	createIndex(db, "table_int", "idx_age_int", "age")

	// Measure time to create index on table_uuid
	createIndex(db, "table_uuid", "idx_age_uuid", "age")
}