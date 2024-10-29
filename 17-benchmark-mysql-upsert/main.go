package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

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

func benchmarkOnDuplicateKeyUpdate(db *sql.DB, n int) {
	for i := 0; i < n; i++ {
		_, err := db.Exec("INSERT INTO test_table (id, data) VALUES (?, ?) ON DUPLICATE KEY UPDATE data = VALUES(data)", i, "test_data")
		if err != nil {
			panic(err)
		}
	}
}

func benchmarkReplaceInto(db *sql.DB, n int) {
	for i := 0; i < n; i++ {
		_, err := db.Exec("REPLACE INTO test_table (id, data) VALUES (?, ?)", i, "test_data")
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	n_arr := []int{1, 10, 100, 500, 1000, 5000, 10000, 50000, 1000000}

	for _, n := range(n_arr){
		// Benchmark ON DUPLICATE KEY UPDATE
		start := time.Now()
		benchmarkOnDuplicateKeyUpdate(db, n)
		fmt.Println("ON DUPLICATE KEY UPDATE  n =", n, "Time Taken:", time.Since(start))

		// Benchmark REPLACE INTO
		start = time.Now()
		benchmarkReplaceInto(db, n)
		fmt.Println("REPLACE INTO\t\t n =", n, "Time Taken:", time.Since(start))
	}
}
