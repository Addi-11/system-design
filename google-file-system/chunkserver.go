package main

import (
	"net/http"
	"log"
)

func (ch *GFSChunkServer) handler(r http.ResponseWriter, w *http.Request){
	
}

func (ch *GFSChunkServer) start(){
	mux := http.NewServeMux()
	http.HandleFunc("/", ch.handler)
	log.Printf("Chunk Server running at %s\n", ch.Address)
	log.Fatal(http.ListenAndServe(ch.Address, mux))
}