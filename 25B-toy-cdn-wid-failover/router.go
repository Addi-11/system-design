package main

import (
	"fmt"
	"net/http"
	"log"
	"sync/atomic"
	"io"
)

type Router struct{
	nodes []*CacheServer
	counter uint32
}

func NewRouter(nodes []*CacheServer) *Router{
	return &Router{
		nodes: nodes,
	}
}

func (router *Router) getNextNode() * CacheServer{
	idx := atomic.AddUint32(&router.counter, 1)%uint32(len(router.nodes))
	return router.nodes[idx]
}

func (router *Router) handler(w http.ResponseWriter, r *http.Request){
	node := router.getNextNode()

	fmt.Printf("Routing request (%s) to CDN: %s\n", r.URL.Path, node.address)

	resp, err := http.Get("http://" + node.address + r.URL.Path)
	if err != nil{
		http.Error(w, "CDN Node error", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}	

func (router *Router) start(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", router.handler)
	fmt.Println("Router running on :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}