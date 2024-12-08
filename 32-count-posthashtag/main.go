package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


const (
	numHashtag int = 6000
	uploads int = 10000
	batchSize int = 1000
)

var DB *sql.DB

func init(){
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"
	dbName := "test_db"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	DB, _ = sql.Open("mysql", dsn)
}


func main(){
	// Create Hashtags Table: uncomment to create a new table
	// insertDBData()

	uploads := generateNewPosts()

	naive(uploads)
	naiveBatching(uploads, batchSize)
	efficientBatchingDeepCopy(uploads, batchSize)
	efficientBatchingTwoMaps(uploads, batchSize)
	kafkaAdapter(uploads)
}

func generateNewPosts() []string{
	var newPosts []string
	for i := 1; i<=uploads; i++{
		newPosts = append(newPosts, fmt.Sprintf("#tag%d", rand.Intn(numHashtag)))
	}
	return newPosts
}

func naive(uploads []string){
	start := time.Now()

	for _, hashtag_id := range(uploads){
		DB.Exec("UPDATE hashtags SET count = count + 1 WHERE hashtag_id = ?", hashtag_id)
	}

	fmt.Printf("Naive Counting, Time Taken: %v\n", time.Since(start))
}

func naiveBatching(uploads []string, batchSize int){
	start := time.Now()

	batch := make(map[string]int)

	for i, hashtag_id := range(uploads){
		batch[hashtag_id]++
		if i%batchSize == 0 || i == len(uploads)-1{
			for tag, cnt := range(batch){
				DB.Exec("UPDATE hashtags SET count = count + ? WHERE hashtag_id = ?", cnt, tag)
			}
			batch = make(map[string]int)
		}
	}

	fmt.Printf("Naive Batching, Time Taken: %v\n", time.Since(start))
	
}

func efficientBatchingDeepCopy(uploads []string, flushInterval int){
	batch := make(map[string]int)
	done := make(chan bool)
	var mu sync.Mutex
	
	// Go routine in parallel for periodic flush, using ticker
	go func(){
		ticker := time.NewTicker(time.Duration(flushInterval) * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <- ticker.C:
				// STOP THE WORLD
				mu.Lock()
					localBatch := make(map[string]int)
					for tag, cnt := range batch {
						localBatch[tag] = cnt
					}
					batch = make(map[string]int)
				mu.Unlock()
				for tag, cnt := range(localBatch){
					DB.Exec("UPDATE hashtags SET count = count + ? WHERE hashtag_id = ?", cnt, tag)
				}
			case <-done:
				return
			}
		}
	}()

	start := time.Now()

	for _, hashtag_id := range(uploads){
		mu.Lock()
		batch[hashtag_id]++ // increment operation is not atomic, so locks
		mu.Unlock()
	}

	close(done)
	fmt.Printf("Efficient Batching (Deep-Copy), Time Taken: %v\n", time.Since(start))
}

func efficientBatchingTwoMaps(uploads []string, flushInterval int){
	activeBatch := make(map[string]int)
	backupBatch := make(map[string]int)
	var mu sync.Mutex
	done := make(chan bool)

	// Go Routine in parallel for periodic flush
	go func(){
		ticker := time.NewTicker(time.Duration(flushInterval) * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// STOP THE WORLD
				mu.Lock()
					current := activeBatch
					activeBatch = backupBatch
					backupBatch = make(map[string]int)
				mu.Unlock()
				for tag, cnt := range(current){
					DB.Exec("UPDATE hashtags SET count = count + ? WHERE hashtag_id = ?", cnt, tag)
				}
			case <-done:
				return
			}
		}
	}()

	start := time.Now()

	for _, hashtag_id := range(uploads){
		mu.Lock()
		activeBatch[hashtag_id]++ // increment operation is not atomic, so locks
		mu.Unlock()
	}

	close(done)
	fmt.Printf("Efficient Batching (Two-Maps), Time Taken: %v\n", time.Since(start))
}

func kafkaAdapter(uploads []string){
	partitions := make(map[string][]string)
	for _, hashtag_id := range(uploads){
		partitions[hashtag_id] = append(partitions[hashtag_id], hashtag_id)
	}

	start := time.Now()

	for tag, events := range(partitions){
		count := len(events)
		DB.Exec("UPDATE hashtags SET count = count + ? WHERE hashtag_id = ?", count, tag)
	}

	fmt.Printf("Kafka Adapter, Time Taken: %v\n", time.Since(start))
}