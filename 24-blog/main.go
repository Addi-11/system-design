package main

import (
	"database/sql"
	"math/rand"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

var dbs []*sql.DB
var mu sync.Mutex
var counter int64
var filename string = "counter.txt"

func init(){
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbHost := "127.0.0.1:3300"

	shard1 := fmt.Sprintf("%s:%s@tcp(%s)/blog_shard1", dbUser, dbPass, dbHost)
	shard2 := fmt.Sprintf("%s:%s@tcp(%s)/blog_shard2", dbUser, dbPass, dbHost)

	db1, _ := sql.Open("mysql", shard1)
	db2, _ := sql.Open("mysql", shard2)

	dbs = append(dbs, db1)
	dbs = append(dbs, db2)
}

func saveCounter(counter int64){
	data := strconv.FormatInt(counter, 10)
	os.WriteFile(filename, []byte(data), 0644)
}

func loadCounter() int64{
	data, err := os.ReadFile(filename)
	if err != nil{
		panic(err)
	}

	id, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil{
		panic(err)
	}
	return id
}

func generateBlogID() int64{
	mu.Lock()
		id := counter
		counter++
		if counter%10 == 0{
			saveCounter(id)
		}
	mu.Unlock()
	return id
}

func getShard(userID int) *sql.DB{
	idx := userID % len(dbs)
	return dbs[idx]
}

func createBlogHandler(w http.ResponseWriter, r *http.Request){
	userID := rand.Intn(10)
	blogID := generateBlogID()
	title := fmt.Sprintf("Blog-%d",blogID)
	content := fmt.Sprintf("This is the content of blog_%d.", blogID)

	db := getShard(userID)
	_, err := db.Exec("INSERT into blogs (blog_id, user_id, title, content) VALUES (?, ?, ?, ?)", blogID, userID, title, content)
	if err != nil{
		http.Error(w, "Failed to insert blog", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Blog created with ID: %d on shard for user %d\n", blogID, userID)
}

func fetchBlogHandler(w http.ResponseWriter, r *http.Request){
	// extract userId from the url
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
	}

	userIDStr := strings.TrimPrefix(parts[2], "userid-")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	db := getShard(userID)
	rows, err := db.Query("SELECT blog_id, title, content FROM blogs WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, "Failed to fetch blogs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// print the values on the webpage
	for rows.Next() {
		var blogID int64
		var title, content string
		if err := rows.Scan(&blogID, &title, &content); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Blog ID: %d, Title: %s, Content: %s\n", blogID, title, content)
	}
}

func main(){
	counter = loadCounter() + 100 //maintain persistence

	http.HandleFunc("/create_blog", createBlogHandler)
	http.HandleFunc("/fetch_blog/", fetchBlogHandler)
	fmt.Println("Starting server at 8080...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}