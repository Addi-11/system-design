package main

import (
	"fmt"
	"time"
)


// lets start server creations
func main(){
	// initialize origin servers
	origin1 := &OriginServer{address: ":8083"}
	origin2 := &OriginServer{address: ":8084"}

	// start the origin servers - on different threads
	go origin1.start()
	go origin2.start()

	// initialize CDN with 2 cache servers
	cdn := []*CacheServer{
		NewCacheServer(":8081", []*OriginServer{origin1, origin2}),
		NewCacheServer(":8082", []*OriginServer{origin1, origin2}),
	}

	// start the CDN servers - on different threads
	for _, c := range(cdn){
		go c.start()
	}

	
	go monitor(cdn, []*OriginServer{origin1, origin2})

	// run router
	router := NewRouter(cdn)
	router.start()


    select {} // run main indefinetly
}

// monitoring function
func monitor(caches []*CacheServer, origins []*OriginServer) {
    ticker := time.NewTicker(10 * time.Second) // Adjust interval as needed
    defer ticker.Stop()

    for range ticker.C {
        fmt.Println("\n----- Monitoring Report -----")
        for _, cache := range caches {
            cache.mu.Lock()
            fmt.Printf("Cache Server %s - Requests served: %d\n", cache.address, cache.load)
            cache.mu.Unlock()
        }

        for _, origin := range origins {
            origin.mu.Lock()
            fmt.Printf("Origin Server %s - Load: %d\n", origin.address, origin.load)
            origin.mu.Unlock()
        }
        fmt.Println("-----------------------------\n")
    }
}