package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"math/rand"
)

type CacheServer struct{
	cache map[string][]byte // req_url, response bytes
	address string
	origins []*OriginServer
	mu sync.RWMutex
	load int
}

func NewCacheServer(address string, origins []*OriginServer) *CacheServer{
	return &CacheServer{
		cache: make(map[string][]byte),
		origins: origins,
		address: address,
	}
}

func (cdn *CacheServer) getOrigin() *OriginServer{
	// we use random-generator to simulate 80/20 load
	if rand.Intn(100)<80{
		return cdn.origins[0]
	}
	return cdn.origins[1]
}

func (cdn *CacheServer) handler(w http.ResponseWriter, r *http.Request){
	// increase load
	cdn.mu.Lock()
		cdn.load ++
	cdn.mu.Unlock()

	// read lock on cdn cache map
	cdn.mu.RLock()
		data, cache := cdn.cache[r.URL.Path]
	cdn.mu.RUnlock()

	// cache hit
	if cache{
		fmt.Printf("Cache hit at %s for %s\n", cdn.address, r.URL.Path)
		w.Write(data)
		return
	}

	// cache miss
	fmt.Printf("Cache miss at %s for %s\n", cdn.address, r.URL.Path)
	origin := cdn.getOrigin()
	fmt.Printf("Fetching from origin %s\n", origin.address)
	resp, err := http.Get("http://" + origin.address + r.URL.Path)
	if err != nil || resp.StatusCode != 200{
		log.Printf("Error fetching from origin %s: %v\n", origin.address, err)
		http.Error(w, "Error fetching from origin", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// read data from origin
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading data", http.StatusInternalServerError)
		return
	}

	// write lock on cdn cache map
	cdn.mu.Lock()
		cdn.cache[r.URL.Path] = data
	cdn.mu.Unlock()

	// write to client
	w.Write(data)
}

func (cdn *CacheServer) start(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", cdn.handler)
	fmt.Println("CDN running on ", cdn.address)
	log.Fatal(http.ListenAndServe(cdn.address, mux))
}