package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
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
	
	w.Header().Set("Content-Type", "text/plain")

	report := router.generateServerHealthReport()

	fmt.Fprint(w, report)
	fmt.Print(report)
}

func (router *Router) generateServerHealthReport() string {
	report := "SERVER HEALTH REPORT\n\n"
	for _, backendServer := range router.config.backendServers {
		report += fmt.Sprintf("Load on origin server %s: %d\n", backendServer.name, backendServer.load)
		report += "Data stored in server: "
		for _, data := range backendServer.data {
			report += fmt.Sprintf("%v, ", data)
		}
		report += "\n\n"
	}
	return report
}

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

func (router *Router) addServerHandler(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("serverName")
	address := fmt.Sprintf(":%s", r.FormValue("serverAddress")) //append : to the port address

	router.addServer(name, address)
	http.Redirect(w, r, "/health", http.StatusSeeOther)
}

func (router *Router) addServer(name string, address string){
	pos := hashFunc(name)
	newServer := &OriginServer{name: name, address: address, pos: pos}

	// add it to the list of backend servers
	idx := sort.Search(len(router.config.backendServers), func(i int) bool {
		return router.config.backendServers[i].pos >= pos
	})

	backendServers := router.config.backendServers
	router.config.backendServers = append(backendServers[:idx], append([]*OriginServer{newServer}, backendServers[idx:]...)...)

	// start the new server
	go newServer.startOriginServer()
	fmt.Printf("Added server %s at position %d\n", newServer.name, newServer.pos)
}

func (router *Router) removeServerHandler(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("removeServerName")
	pos := hashFunc(name)

	// index of the server to remove
	idx := sort.Search(len(router.config.backendServers), func(i int) bool {
		return router.config.backendServers[i].pos >= pos
	})
	oldServer :=  router.config.backendServers[idx]

	// migrate data from the server to the next server
	nextIdx := (idx + 1) % len(router.config.backendServers)
	nextServer := router.config.backendServers[nextIdx]
	nextServer.data = append(nextServer.data, oldServer.data...)

	oldServer.data = nil// clear data of the current idx

	fmt.Printf("Migrating data from %s server to %s server\n", oldServer.name, nextServer.name)

	// remove the server by excluding it from the list
	router.config.backendServers = append(router.config.backendServers[:idx], router.config.backendServers[idx+1:]...)
	fmt.Printf("Removed server %s at position %d\n", name, pos)


	http.Redirect(w, r, "/health", http.StatusSeeOther)
}

func (router *Router) startRouter(){
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", router.formHandler) // serve  the HTML form
	mux.HandleFunc("/submit", router.handler) // handle form submission
	mux.HandleFunc("/health", router.healthCheckHandler)
	mux.HandleFunc("/addServer", router.addServerHandler)
	mux.HandleFunc("/removeServer", router.removeServerHandler)
	fmt.Printf("Starting router server at port %s\n", router.address)
	log.Fatal(http.ListenAndServe(router.address, mux))
}