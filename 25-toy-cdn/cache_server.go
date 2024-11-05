package main

import (
	"fmt"
	"net/http"
	"log"
	"io"
	"sync"
)

type CacheNode struct{
	cache map[string][]byte //string, bytes
	mu sync.RWMutex
	origin string
	address string
}

func NewCacheNode(address, origin string) *CacheNode{
	return &CacheNode{
		cache: make(map[string][]byte),
		origin: origin,
		address: address,
	}
}

func (c *CacheNode) cacheHandler(w http.ResponseWriter, r *http.Request){
	c.mu.RLock()
	data, cached := c.cache[r.URL.Path]
	c.mu.RUnlock()

	if cached{
		fmt.Printf("Cache hit at %s for %s.\n", c.address, r.URL.Path)
		w.Write(data)
		return
	}

	// cache miss
	fmt.Printf("Cache miss at %s for %s.\n", c.address, r.URL.Path)
	resp, err := http.Get(c.origin + r.URL.Path)
	if err != nil || resp.StatusCode != 200{
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

	// cache the data
	c.mu.Lock()
	c.cache[r.URL.Path] = data
	c.mu.Unlock()

	// serve to client
	w.Write(data)
}

func (c *CacheNode) cacheServerStart(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", c.cacheHandler)
	fmt.Println("Cache node running on ", c.address)
	log.Fatal(http.ListenAndServe(c.address, mux))
}