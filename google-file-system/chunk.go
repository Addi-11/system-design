package main

import (
	"fmt"
	"log"
	"net/http"
)

type ChunkServer struct {
	Address string
	Chunks  map[string]string // ChunkID -> Data
}

func (cs *ChunkServer) WriteChunkHandler(w http.ResponseWriter, r *http.Request) {
	chunkID := r.URL.Query().Get("chunkID")
	data := r.URL.Query().Get("data")
	cs.Chunks[chunkID] = data
	fmt.Fprintf(w, "Chunk %s written successfully.\n", chunkID)
}

func (cs *ChunkServer) ReadChunkHandler(w http.ResponseWriter, r *http.Request) {
	chunkID := r.URL.Query().Get("chunkID")
	data, exists := cs.Chunks[chunkID]
	if !exists {
		http.Error(w, "Chunk not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, data)
}

func (cs *ChunkServer) Start() {
	http.HandleFunc("/write_chunk", cs.WriteChunkHandler)
	http.HandleFunc("/read_chunk", cs.ReadChunkHandler)
	log.Printf("ChunkServer running at %s\n", cs.Address)
	log.Fatal(http.ListenAndServe(cs.Address, nil))
}
