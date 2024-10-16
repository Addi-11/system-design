package main

import (
	"fmt"
	"sync"
	"time"
	"database/sql"
	"os"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser = os.Getenv("DBUSER")
	dbPass = os.Getenv("DBPASS")
	dbHost = "127.0.0.1:3300"
	dbName = "test_db"

	// connection string to the DB
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
)

// open connection to the DB
func createNewConnection() (*sql.DB, error){
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	return db, nil
}

// simulate database operation
func simulateDBOperation (db *sql.DB) error {
	_, err := db.Exec("SELECT SLEEP(0.01);")
	return err
}

func benchmarkNonPool(n int) (time.Duration, error){
	start := time.Now()
	var wg sync.WaitGroup
	var once sync.Once
	var operr error

	helper := func (){
		defer wg.Done()
		conn, err := createNewConnection()
		if err != nil {
			once.Do(func() {operr = err})
			return
		}
		defer conn.Close()
		if err := simulateDBOperation(conn); err != nil {
			once.Do(func(){
				operr = err
			})
		}
	}

	for i:= 0; i < n; i++ {
		wg.Add(1)
		go helper()
	}
	// wait for all goroutines to complete
	wg.Wait()
	return time.Since(start), operr
}


// Conection pool struct to manage DB connections
type ConnectionPool struct {
	mu sync.Mutex
	channel chan interface{} // channel represent s blocking queue
	conns []*sql.DB
}


func NewPool(maxConn int)(*ConnectionPool, error){
	pool := &ConnectionPool{
		conns		: make([]*sql.DB, 0, maxConn),
		channel 	: make(chan interface{}, maxConn),
	}

	for i:=0; i<maxConn; i++{
		conn, err := createNewConnection()
		if err != nil {
			return nil, err
		}
		pool.conns = append(pool.conns, conn)
		pool.channel <- nil // indicates connection is ready
	}

	return pool, nil
}

func (pool *ConnectionPool) Get()(*sql.DB, error){
	// Block until there is an available connection
	<-pool.channel

	pool.mu.Lock()
	defer pool.mu.Unlock()

	conn :=  pool.conns[0] // Get the connection
	pool.conns = pool.conns[1:] // Remove connection from the pool
	return conn, nil
}

func (pool *ConnectionPool) Put(conn *sql.DB) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	pool.conns = append(pool.conns, conn)

	// signal one connection is available
	pool.channel <- nil
}

func (pool *ConnectionPool) Close() {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	for _, conn := range pool.conns{
		conn.Close()
	}

	pool.conns = nil
}

func benchmarkPool(n int) (time.Duration, error){
	// creating a connection pool
	pool, err := NewPool(100)
	if err != nil{
		return 0, fmt.Errorf("Error intializing connection pool: %w", err)
	}

	start := time.Now()
	var wg sync.WaitGroup
	var once sync.Once
	var operr error

	helper := func(){
		defer wg.Done()
		// get thread from the connection pool
		conn, err := pool.Get()
		if err != nil {
			once.Do(func(){
				operr = err
			})
		}
		if err := simulateDBOperation(conn); err != nil {
			once.Do(func(){
				operr = err
			})
		}

		// return thread to the connection pool
		pool.Put(conn)
	}

	for i:=0; i<n; i++ {
		wg.Add(1)
		go helper()
	}

	wg.Wait()
	// close the connection pool
	pool.Close()
	return time.Since(start), operr
}

func main(){
	tests := []int{10, 100, 200, 300, 500, 1000}
	
	fmt.Println("\nStarting non-pool benchmarks:")

	for _, n := range tests {
		elapsed, err := benchmarkNonPool(n)

		if err != nil {
			fmt.Printf("Error running non-pool benchmark for %d threads: %v\n", n, err)
		} else {
			fmt.Printf("Non-pool time for %d threads: %v\n", n, elapsed)
		}
	}

	fmt.Println("\nStarting pool benchmarks:")

	for _, n := range tests {
		elapsed, err := benchmarkPool(n)
		if err != nil {
			fmt.Printf("Error running pool benchmark for %d threads: %v\n", n, err)
		} else {
			fmt.Printf("Pool time for %d threads: %v\n", n, elapsed)
		}
	}
	
}