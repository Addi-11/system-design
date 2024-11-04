package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "time"
	"os"
)

func main() {
    startTime := time.Now()

	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"

	sourceDSN := fmt.Sprintf("%s:%s@tcp(%s)/insta1", dbUser, dbPass, dbHost)
	targetDSN := fmt.Sprintf("%s:%s@tcp(%s)/insta2", dbUser, dbPass, dbHost)

	sourceDB, err := sql.Open("mysql", sourceDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceDB.Close()

	targetDB, err := sql.Open("mysql", targetDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer targetDB.Close()

    // query rows from insta1.posts
    rows, err := sourceDB.Query("SELECT post_id, user_id, content, created_at FROM posts")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // iterate through rows and insert into insta2.posts
    var rowCount int
    for rows.Next() {
        rowStartTime := time.Now()

        var postID int
        var userID int
        var content string
        var createdAt string

        if err := rows.Scan(&postID, &userID, &content, &createdAt); err != nil {
            log.Fatal(err)
        }

        _, err = targetDB.Exec("INSERT INTO posts (post_id, user_id, content, created_at) VALUES (?, ?, ?, ?)", postID, userID, content, createdAt)
        if err != nil {
            log.Fatal(err)
        }

        rowDuration := time.Since(rowStartTime)
        fmt.Printf("Migrated post ID %d (Row time: %v)\n", postID, rowDuration)

        rowCount++
    }

    totalDuration := time.Since(startTime)
    fmt.Printf("\nMigration completed.\nTotal rows migrated: %d\nTotal migration time: %v\n", rowCount, totalDuration)
}
