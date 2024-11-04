package main

import (
	"database/sql"
	"fmt"
	"log"
	// "math/rand"
	// "strings"
	// "time"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func insert_main() {
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

	insertData(sourceDB)
	insertData(targetDB)

	fmt.Println("Data insertion complete.")
}

func insertData(db *sql.DB) {
	const numRecords = 1000
	for i := 1; i <= numRecords; i++ {
		username := fmt.Sprintf("user%d", i)
		email := fmt.Sprintf("user%d@example.com", i)

		// insert into users table
		_, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
		if err != nil {
			log.Fatalf("Error inserting into users: %v", err)
		}

		// user_id is auto-incremented - use it for posts and profiles
		var userID int
		err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&userID)
		if err != nil {
			log.Fatalf("Error fetching last inserted user_id: %v", err)
		}

		// insert into posts table
		content := fmt.Sprintf("This is the content of post %d", i)
		_, err = db.Exec("INSERT INTO posts (user_id, content) VALUES (?, ?)", userID, content)
		if err != nil {
			log.Fatalf("Error inserting into posts: %v", err)
		}

		// insert into profile table
		bio := fmt.Sprintf("This is the bio for user %d", i)
		_, err = db.Exec("INSERT INTO profile (user_id, bio) VALUES (?, ?)", userID, bio)
		if err != nil {
			log.Fatalf("Error inserting into profile: %v", err)
		}
	}

	fmt.Printf("Inserted %d records into the database.\n", numRecords)
}
