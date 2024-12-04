package main

import "time"

func main() {
	// Start the chunk server
	chunkServer := &ChunkServer{Address: ":8081", Chunks: make(map[string]string)}
	go chunkServer.Start()

	// Start the master server
	masterServer := &MasterServer{
		Address:       ":8080",
		FileToChunks:  make(map[string][]string),
		ChunkToServer: make(map[string]string),
	}
	go masterServer.Start()

	// Wait for the servers to start
	time.Sleep(2 * time.Second)

	// Client operations
	client := &Client{MasterAddress: ":8080"}
	client.WriteFile("file1.txt", "Hello, GFS!")
	client.ReadFile("file1.txt")

	select{}
}
