package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// server sent event handler
func SSEHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// sent data to the client
	for {
		fmt.Fprintf(w, "data: Current time is %s\n\n", time.Now().Format(time.RFC3339))

		// flush the data to the client - immideately, not buffer
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}else {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func main(){
	// register the SSE handler
	http.HandleFunc("/events", SSEHandler)

	// start the server
	log.Println("starting server on: 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}