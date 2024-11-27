package main

type GFSClient struct {
	MasterAdress string
}

type GFSMaster struct {
	Files map[string][]string // file -> chunk mapping: files are broken to chunks
	ChunkLocation map[string][]string // data-> chunkserver mapping: servers where is the chunkdata present
	Servers map[string]GFSChunkServer
}

type GFSChunkServer struct {
	ChunkID string
	Address string
	Chunks map[string]Chunk
}

type Chunk struct {
	ID string
	Data []byte
}