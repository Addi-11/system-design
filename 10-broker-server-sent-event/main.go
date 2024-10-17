package main

import (
	"fmt"
	"net/http"
	"time"
	"sync"
	"log"
)

// create a broker
type Broker struct{
	Notifier chan []byte // receives events from outside, broadcasts them to clients
	newClients chan chan []byte // channel to register the client, and create a new channel for the new client
	closingClients chan chan []byte //
	clients map[chan []byte]bool
	mu sync.Mutex // protect the client map
}

// creating a new server
func NewServer() *Broker{
	broker := &Broker{
		Notifier: make(chan []byte, 1),
		newClients: make(chan chan[] byte),
		closingClients: make(chan chan[]byte),
		clients: make(map[chan []byte]bool),
	}

	go broker.listen()
	return broker
}

func (broker *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Notify the new connection with a welcome message
	fmt.Fprintf(w, "Welcome to the SSE server!\n\n")
	flusher, ok := w.(http.Flusher)
	if !ok{
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// add a new channel for a new client
	msgChannel := make(chan []byte)
	broker.newClients <- msgChannel

	// close the client - when the browser closes, or client navigates away from the page
	notify := r.Context().Done()
	go func(){
		<- notify
		broker.closingClients <- msgChannel
	}()

	defer func(){
		broker.closingClients <- msgChannel
	}()

	// flush text to the events
	for {
		fmt.Fprintf(w, "data: %s\n\n", <-msgChannel)
		flusher.Flush()
	}
}

func (broker *Broker) listen(){
	for {
		select {
		case s := <-broker.newClients:
			broker.mu.Lock()
			broker.clients[s] = true
			broker.mu.Unlock()
			log.Printf("Client added. %d registered clients.", len(broker.clients))

		case s := <-broker.closingClients:
			broker.mu.Lock()
			delete(broker.clients, s)
			broker.mu.Unlock()
			log.Printf("Removed client. %d registered clients.", len(broker.clients))
		
		case event := <-broker.Notifier:
			broker.mu.Lock()
			for clientMessageChan := range broker.clients {
				select {
				case clientMessageChan <-event:
					log.Println("Message sent to a client.", clientMessageChan)
				default:
					log.Println("Client too slow, skipping message.", clientMessageChan)
				}
			}
			broker.mu.Unlock()

		}
	}
}

func main(){
	broker := NewServer()

	go func() {
		for {
			time.Sleep(time.Second * 5)
			eventString := fmt.Sprintf("the time is %v", time.Now())
			log.Println("Broadcasting event")
			broker.Notifier <- []byte(eventString)
		}
	}()

	log.Fatal(http.ListenAndServe(":3000", broker))
}