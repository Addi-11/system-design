package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func originHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Serving file from origin: ", r.URL.Path)
	time.Sleep(5 * time.Second) //simulate latency
	http.ServeFile(w, r, "./static" + r.URL.Path)
}

func originServerStart(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", originHandler)
	fmt.Println("Origin Server running on : 8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}