package main

import (
	"fmt"
	"net/http"
	"log"
)

type OriginServer struct{
	name string
	address string
	load int
	data []interface{} // store key data in the origin server
}

func (origin *OriginServer) handler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	key := r.FormValue("key")
	origin.data = append(origin.data, key)

	fmt.Printf("Storing data in origin server %s\nCurrentLoad: %d\n", origin.name, origin.load)
}

func (origin *OriginServer) startOriginServer(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", origin.handler)
	fmt.Printf("Starting origin server %s at port %s\n", origin.name, origin.address)
	log.Fatal(http.ListenAndServe(origin.address, mux))
}