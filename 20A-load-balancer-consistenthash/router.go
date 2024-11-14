package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
)

type RouterConfig struct{
	algo string
	backendServers []*OriginServer // backend servers stored in sorted order
}

type Router struct{
	address string
	config RouterConfig
	next int // counter for round-robin
}

func (router *Router) formHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html")
}

func (router *Router) handler(w http.ResponseWriter, r *http.Request){
	// get key and values forms
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	key := r.FormValue("key")

	// get origin server
	origin := router.getNextServer(key)
	originAddress := fmt.Sprintf("http://localhost%s", origin.address)

	// forward request to the origin server
	resp, err := http.PostForm(originAddress, map[string][]string{"key": {key}})
	if err != nil{
		http.Error(w, "Failed to reach origin server", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	origin.load++
	http.Redirect(w, r, "/", http.StatusSeeOther) // Redirect back to the form
}

func (router *Router) healthCheckHandler(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	for _, backendServer := range(router.config.backendServers){
		fmt.Printf("Load on origin server %s: %d\n", backendServer.name, backendServer.load) // print in the log
		fmt.Fprintf(w, "Load on origin server %s: %d\n", backendServer.name, backendServer.load) // print on the html page
	}
}

// func (router *Router) addServer() *OriginServer{

// }

func (router *Router) getNextServer(key string) *OriginServer{
	var origin *OriginServer

	switch router.config.algo {
	case "round_robin":
		origin = router.config.backendServers[router.next]
		router.next = (router.next + 1) % len(router.config.backendServers)
	case "random_selection":
		idx := rand.Intn(len(router.config.backendServers))
		origin = router.config.backendServers[idx]
	case "consistent_hash":
		origin = getServerConsistentHash(key, router.config.backendServers) // we get this based on the input key
	}

	return origin
}

func (router *Router) startRouter(){
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", router.formHandler) // serve  the HTML form
	mux.HandleFunc("/submit", router.handler) // handle form submission
	mux.HandleFunc("/health", router.healthCheckHandler)
	fmt.Printf("Starting router server at port %s\n", router.address)
	log.Fatal(http.ListenAndServe(router.address, mux))
}