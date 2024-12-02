package main

import (
	"log"
	"net/http"
)

func (master *GFSMaster) handler(r http.ResponseWriter, w *http.Request){
	
}

func (master *GFSMaster)start(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", master.handler)
	log.Printf("GFS Master running at %s\n", master.Address)
	log.Fatal(http.ListenAndServe(master.Address, mux))
}