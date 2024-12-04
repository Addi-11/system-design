package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MasterServer struct {
	Address       string
	FileToChunks  map[string][]string // FilePath -> List of ChunkIDs
	ChunkToServer map[string]string   // ChunkID -> ChunkServer Address
}

func (ms *MasterServer) WriteFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("filePath")
	data := r.URL.Query().Get("data")
	chunkID := fmt.Sprintf("chunk-%d", len(ms.FileToChunks)+1)
	ms.FileToChunks[filePath] = []string{chunkID}
	ms.ChunkToServer[chunkID] = ":8081" // Assign to ChunkServer 1 (hardcoded for simplicity)
	resp := map[string]string{"chunkID": chunkID, "chunkServer": ms.ChunkToServer[chunkID]}
	json.NewEncoder(w).Encode(resp)
}

func (ms *MasterServer) ReadFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("filePath")
	chunks, exists := ms.FileToChunks[filePath]
	if !exists {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	resp := map[string]string{"chunkID": chunks[0], "chunkServer": ms.ChunkToServer[chunks[0]]}
	json.NewEncoder(w).Encode(resp)
}

func (ms *MasterServer) Start() {
	http.HandleFunc("/write_file", ms.WriteFileHandler)
	http.HandleFunc("/read_file", ms.ReadFileHandler)
	log.Printf("MasterServer running at %s\n", ms.Address)
	log.Fatal(http.ListenAndServe(ms.Address, nil))
}
