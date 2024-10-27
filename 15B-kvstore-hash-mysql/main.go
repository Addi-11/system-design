package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB []*sql.DB
const shardCount = 2

func failOnError(err error, msg string){
	if err !=nil{
		log.Fatalf("%s: %s", err, msg)
	}
}

func init(){
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "test_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	db1, err := sql.Open("mysql", dsn)
	failOnError(err, "Error connecting to DB1")

	err = db1.Ping()
	failOnError(err, "Error pinging DB1")

	DB = append(DB, db1)

	db2, err := sql.Open("mysql", dsn)
	failOnError(err, "Error connecting to DB2")

	err = db2.Ping()
	failOnError(err, "Error pinging DB2")

	DB = append(DB, db2)
}

func getShard(key string) (*sql.DB, int){
	hash := sha256.Sum256([]byte(key))
	idx := int(hash[0]) % shardCount
	return DB[idx], idx+1
}

func put(key string, value string, ttl time.Duration) {
    db, idx := getShard(key)
    table := fmt.Sprintf("kv_store_shard%d", idx)
    expiry := time.Now().Add(ttl)

	formattedExpiry := expiry.Format("2006-01-02 15:04:05") // format it before storing - else timezone may change

    _, err := db.Exec(fmt.Sprintf("REPLACE INTO %s (`key`, `value`, `expired_at`) VALUES (?, ?, ?)", table), key, value, formattedExpiry)
    failOnError(err, "Error putting value")
}




func get(key string) string {
	db, idx := getShard(key)
	table := fmt.Sprintf("kv_store_shard%d", idx)
	var value string
	var expiredAt time.Time

	err := db.QueryRow(fmt.Sprintf("SELECT value from %s WHERE `key` = ? AND expired_at > NOW()", table), key).Scan(&value)

	if err == sql.ErrNoRows {
		fmt.Printf("Key not found or expired: %s\n", key)
		return ""
	} else if err != nil {
		log.Fatalf("Get Failed: %v\n", err)
	}

	fmt.Printf("Retrieved key '%s' with value '%s', expires at %s\n", key, value, expiredAt)
	return value
}


func del(key string){
	db, idx := getShard(key)
	table := fmt.Sprintf("kv_store_shard%d", idx)

	_, err := db.Exec(fmt.Sprintf("UPDATE %s SET expired_at = '1970-01-01' WHERE `key` = ? AND expired_at > NOW()", table), key)
	failOnError(err, "Failed to delete key")
}

func ttl(){
	for i, db := range(DB){
		_, err := db.Exec(fmt.Sprintf("DELETE FROM kv_store_shard%d WHERE expired_at <= NOW() LIMIT 1000", i+1))
		failOnError(err, "Failed cleaning the expired keys.")
	}
}

func main(){
	num := 20

	for i := 0; i < num; i++{
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", rand.Intn(100))
		ttl := time.Duration(3+rand.Intn(10)) * time.Minute

		// Put into table
		put(key, value, ttl)
		fmt.Printf("Put: %s -> %s (TTL: %s)\n", key, value, ttl)

		// randomly get and del keys
		if rand.Intn(100) < 40{
			val := get(key)
			fmt.Printf("Get: %s ->%s\n", key, val)
			time.Sleep(3 * time.Second)
			
			del(key)
			fmt.Printf("Deleted: %s\n", key)
		}
	}
	fmt.Println("Cleaning Expired keys....")
	time.Sleep(50 * time.Second)
	ttl()
	fmt.Println("Expired keys cleanup completed.\n")
}