package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	NumRecords = 1000
	Limit      = 50
	TotalPages = NumRecords / Limit
)

var DatabaseDSN string  = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3300)/insta1", os.Getenv("DBUSER"), os.Getenv("DBPASS"))


func benchmark_main() {
	db, err := sql.Open("mysql", DatabaseDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Benchmark Limit-Offset Pagination
	fmt.Println("Benchmarking Limit-Offset Pagination...")
	start := time.Now()
	benchmarkLimitOffset(db)
	duration := time.Since(start)
	fmt.Printf("Total Time for Limit-Offset: %v\n", duration)

	// Benchmark ID Limit Pagination
	fmt.Println("Benchmarking ID Limit Pagination...")
	start = time.Now()
	benchmarkIDLimit(db)
	duration = time.Since(start)
	fmt.Printf("Total Time for ID Limit: %v\n", duration)
}

// Benchmark Limit-Offset Pagination
func benchmarkLimitOffset(db *sql.DB) {
	for i := 0; i < TotalPages; i++ {
		offset := i * Limit
		rows, err := db.Query("SELECT * FROM users LIMIT ? OFFSET ?", Limit, offset)
		if err != nil {
			log.Fatalf("Query failed: %v", err)
		}
		defer rows.Close()

		// Simulate processing of rows
		for rows.Next() {
			// Assume a simple structure, e.g., int id, string name
			var id int
			var name string
			var email string
			if err := rows.Scan(&id, &name, &email); err != nil {
				log.Fatal(err)
			}
		}
	}
}

// Benchmark ID Limit Pagination
func benchmarkIDLimit(db *sql.DB) {
	var lastID int
	for i := 0; i < TotalPages; i++ {
		rows, err := db.Query("SELECT * FROM users WHERE user_id > ? ORDER BY user_id LIMIT ?", lastID, Limit)
		if err != nil {
			log.Fatalf("Query failed: %v", err)
		}
		defer rows.Close()

		// Simulate processing of rows
		for rows.Next() {
			var id int
			var name string
			var email string
			if err := rows.Scan(&id, &name, &email); err != nil {
				log.Fatal(err)
			}
			lastID = id // Update lastID for the next iteration
		}
	}
}
