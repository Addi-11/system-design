package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (client *GFSClient) WriteFile(filepath string){
	fmt.Printf("Writing file %s to GFS...\n", filepath)
	data, err := os.ReadFile(filepath)
	errorlog("Error reading input file", err)

	// split data into chunks
	chunks := splitDataToChunks(data)

	// contact master to allocate the chunks
	chunkIDs, err := 

	// send each chunk to the chunkserver

}

func (client *GFSClient) ReadFile(filepath string){
	fmt.Printf("Reading file %s to GFS...\n", filepath)

}

func (client *GFSClient) WriteFileHandler(w http.ResponseWriter, r *http.Request){
	filePath := r.URL.Query().Get("filepath")
	client.WriteFile(filePath)
	fmt.Fprintf(w, "File %s written successfully.", filePath)
}

func (client *GFSClient) ReadFileHandler(w http.ResponseWriter, r *http.Request){
	filePath := r.URL.Query().Get("filepath")
	client.ReadFile(filePath)
}

func (client *GFSClient) start(){
	mux := http.NewServeMux()
	mux.HandleFunc("/write_file", client.WriteFileHandler)
	mux.HandleFunc("/read_file", client.ReadFileHandler)
	log.Printf("GFS Client running at %s\n", client.Address)
	log.Fatal(http.ListenAndServe(":8080", mux))
}