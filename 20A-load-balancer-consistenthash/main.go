package main

func main(){
	// intialize 5 origin servers
	originServers := []*OriginServer{
		&OriginServer{name: "A", address: ":8081"},
		&OriginServer{name: "B", address: ":8082"},
		&OriginServer{name: "C", address: ":8083"},
		&OriginServer{name: "D", address: ":8084"},
		&OriginServer{name: "E", address: ":8085"},
	}

	// starting the origin servers - on different threads
	for _, origin := range(originServers){
		go origin.startOriginServer()
	}

	// intializing router config
	routerConfig := &RouterConfig{
		algo: "consistent_hash",
		backendServers: originServers,
	}

	// intializing router
	router := &Router{
		address: ":8080",
		config: *routerConfig,
	}

	// starting router server	
	go router.startRouter()

	select {} //prevent main from exit, before the threads starts
}