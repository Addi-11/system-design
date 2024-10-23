package main

import (
	"fmt"
	"database/sql"
	"os"
	"log"
	"time"

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
	err := DB.QueryRow("SELECT * FROM kv_store WHERE key=? AND expired_at > NOW()", key).Scan(&value)
	if err != nil{
		return "", err
	}
	return value, err
}

func del(key string) error{
	_, err := DB.Exec("UPDATE kv_store SET expired_at = -1 WHERE key = ? AND expired_at >= NOW()", key)
	return err
}

func ttl() error{
	_, err := DB.Exec("DELETE FROM kv_store WHERE expired_at <= NOW()")
	return err
}

func main(){
	put("123", "123abcd", time.Second *20)
	put("24", "24shimla", time.Second *120)
	put("67", "67kolkata", time.Second *220)
}