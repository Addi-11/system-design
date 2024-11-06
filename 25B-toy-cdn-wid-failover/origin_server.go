package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"os"
)

type OriginServer struct{
	address string
	load int
	mu sync.Mutex
}

func (origin *OriginServer) handler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Serving file from origin %s for %s\n", origin.address, r.URL.Path)
	// increase load on the current origin
	origin.mu.Lock()
		origin.load++		
		fmt.Printf("Current load on origin %s: %d\n", origin.address, origin.load)
	origin.mu.Unlock()

	// simulate latency
	time.Sleep(1 * time.Second) 

	// serving files from static folder
	filePath := "./static" + r.URL.Path
	if _, err := os.Stat(filePath); err == nil {
        http.ServeFile(w, r, filePath)
    } else {
        http.Error(w, "File not found", http.StatusNotFound)
    }
}

func (origin *OriginServer) start(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", origin.handler)
	fmt.Println("Running origin server on", origin.address)
	log.Fatal(http.ListenAndServe(origin.address,  mux))
}