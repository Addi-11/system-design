package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init(){
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "test_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil{
		panic(err)
	}

	err = DB.Ping()
	if err != nil{
		log.Fatal("Error pinging DB: ", err)
	}
}

func put(key string, value string, ttl time.Duration) error{
	expired_at := time.Now().Add(ttl)
	_, err := DB.Exec("REPLACE INTO kv_store VALUES (?,?,?)", key, value, expired_at)
	return err
}

func get(key string) (string, error){
	var value string
	err := DB.QueryRow("SELECT value FROM kv_store WHERE `key` = ? AND expired_at > NOW()", key).Scan(&value)
	if err != nil{
		return "", err
	}
	return value, err
}

func del(key string) error{
	_, err := DB.Exec("UPDATE kv_store SET expired_at = '1970-01-01' WHERE `key` = ? AND expired_at > NOW()", key)
	return err
}

func ttl() error{
	_, err := DB.Exec("DELETE FROM kv_store WHERE expired_at <= NOW() LIMIT 1000")
	return err
}

func cli(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		commands := strings.Split(input, " ")

		switch commands[0]{

		case "put":
			if len(commands) != 4 {
				fmt.Println("Usage: put <key> <value> <ttl_in_seconds>")
				continue
			}
			key := commands[1]
			value := commands[2]
			ttl, err := time.ParseDuration(commands[3] + "s")
			if err != nil {
				fmt.Println("Invalid TTL format. Use an integer for seconds.")
				continue
			}
			err = put(key, value, ttl)
			if err != nil {
				fmt.Println("Error putting value:", err)
			} else {
				fmt.Println("Put operation successful")
			}

		case "get":
			if len(commands) != 2 {
				fmt.Println("Usage: get <key>")
				continue
			}
			key := commands[1]
			value, err := get(key)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Println("Key not found or expired")
				} else {
					fmt.Println("Error getting value:", err)
				}
			} else {
				fmt.Println("Value:", value)
			}

		
		case "del":
			if len(commands) != 2 {
				fmt.Println("Usage: del <key>")
				continue
			}
			key := commands[1]
			err := del(key)
			if err != nil {
				fmt.Println("Error deleting key:", err)
			} else {
				fmt.Println("Delete operation successful")
			}

		case "ttl":
			err := ttl()
			if err != nil {
				fmt.Println("Error running TTL:", err)
			} else {
				fmt.Println("TTL cleanup successful")
			}

		case "exit":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Unknown command. Supported commands: put, get, del, ttl, exit")
		}
	}
}

func ttlCleanupRoutine() {
	ticker := time.NewTicker(120 * time.Second) // run every 2 min
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := ttl()
			if err != nil {
				fmt.Println("Error running automatic TTL cleanup:", err)
			} else {
				fmt.Println("Automatic TTL cleanup successful")
			}
		}
	}
}


func main(){
	go ttlCleanupRoutine()
	cli()
}