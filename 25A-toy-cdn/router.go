package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"io"
)

type RoundRobinRouter struct {
	nodes []*CacheNode
	counter uint32
}

func NewRouter(nodes []*CacheNode) *RoundRobinRouter{
	return &RoundRobinRouter{nodes: nodes}
}

func (r *RoundRobinRouter) getNextNode() *CacheNode{
	node := r.nodes[atomic.AddUint32(&r.counter, 1)%uint32(len(r.nodes))]
	return node
}

func (r *RoundRobinRouter) routerHandler(w http.ResponseWriter, rq *http.Request){
	node := r.getNextNode()
	fmt.Printf("Routing request %s to node %s\n", rq.URL.Path, node.address)

	resp, err := http.Get("http://" + node.address + rq.URL.Path)
	if err != nil {
		http.Error(w, "Cache node error", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// copy response to client
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (r *RoundRobinRouter) routerServerStart(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.routerHandler)
	fmt.Println("Router Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}

