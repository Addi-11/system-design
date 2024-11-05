package main

func main(){
	// start origin server
	origin := "http://localhost:8081/"
	go originServerStart()

	// create cached nodes
	cachedNodes := []*CacheNode{
		NewCacheNode(":8082", origin),
		NewCacheNode(":8083", origin),
	}

	for _, node := range(cachedNodes){
		go node.cacheServerStart()
	}

	// start router
	router := NewRouter(cachedNodes)
	router.routerServerStart()
}