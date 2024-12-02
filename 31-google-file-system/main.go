package main

type GFSClient struct {
	Address string
	GFSMasterAddress string
}

type GFSMaster struct {
	Address        string                // Master server address
	Files          map[string][]string   // File-to-chunk mapping
	ChunkLocation  map[string][]string   // DataChunk-to-Chunkserver mapping
	ChunkVersions  map[string]int        // Chunk versioning to ensure consistency
	Namespace      map[string]struct{}   // File namespace metadata (file/directory hierarchy)
}

type GFSChunkServer struct {
	Address string                  // Chunk server address
	Chunks  map[string]Chunk        // ChunkID -> Chunk (chunks stored on this server)
}

type Chunk struct {
	ID       string
	Version  int
	Data     []byte
}


func main(){
	// intialize 3 chunkservers
	chunkServer1 := &GFSChunkServer{Address: ":8081", Chunks: make(map[string]Chunk)}
	chunkServer2 := &GFSChunkServer{Address: ":8082", Chunks: make(map[string]Chunk)}
	chunkServer3 := &GFSChunkServer{Address: ":8083", Chunks: make(map[string]Chunk)}

	// start the servers
	go chunkServer1.start()
	go chunkServer2.start()
	go chunkServer3.start()

	// intialize the master server
	gfsMaster := &GFSMaster{
		Address:       ":8084",
		Files:         make(map[string][]string),
		ChunkLocation: make(map[string][]string),
		ChunkVersions: make(map[string]int),
		Namespace:     make(map[string]struct{}),
	}
	go gfsMaster.start()

	// intialize the gfs client
	gfsClient := &GFSClient{Address: ":8080", GFSMasterAddress: gfsMaster.Address}
	go gfsClient.start()

	select {} // keep the main function running
}