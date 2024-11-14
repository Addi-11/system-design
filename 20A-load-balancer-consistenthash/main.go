package main

func main(){

	// intializing router config
	routerConfig := &RouterConfig{
		algo: "consistent_hash",
	}

	// intializing router
	router := &Router{
		address: ":8080",
		config: *routerConfig,
	}
	
	// intialize 4 origin servers
	originServers := []*OriginServer{
		&OriginServer{name: "A", address: ":8081"},
		&OriginServer{name: "B", address: ":8082"},
		&OriginServer{name: "C", address: ":8083"},
		&OriginServer{name: "D", address: ":8084"},
	}

	// starting the origin servers - on different threads
	for _, origin := range(originServers){
		router.addServer(origin.name, origin.address) // add origin servers to the router config, in sorted order
		go origin.startOriginServer()
	}

	// starting router server	
	go router.startRouter()

	select {} //prevent main from exit, before the threads starts
}