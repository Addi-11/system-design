package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

type BackendServer struct {
	ID   int
	Load int    // Simulating load on the current backend server
	Port string // Port on which the backend server listens
}

type LoadBalancer struct {
	port   string
	config *Config
	mu     sync.Mutex
	next   int // For round-robin
}

// Load balancer config
type Config struct {
	BackendServers  []*BackendServer // List of backend servers
	HealthEndpoint  string           // Endpoint to check the health of the load balancer
	RoutingAlgo     string           // Type of routing algorithm - round_robin, random_selection
}

func newLoadBalancer() *LoadBalancer {
	servers := []*BackendServer{
		{ID: 1, Load: 0, Port: "8001"},
		{ID: 2, Load: 0, Port: "8002"},
		{ID: 3, Load: 0, Port: "8003"},
		{ID: 4, Load: 0, Port: "8004"},
	}
	return &LoadBalancer{
		port:   "8000",
		config: &Config{BackendServers: servers, HealthEndpoint: "/health", RoutingAlgo: "round_robin"},
		next:   0,
	}
}

// Get the next backend server based on the routing algorithm
func (lb *LoadBalancer) nextRequest() *BackendServer {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	var server *BackendServer

	switch lb.config.RoutingAlgo {
	case "round_robin":
		server = lb.config.BackendServers[lb.next]
		lb.next = (lb.next + 1) % len(lb.config.BackendServers)
	case "random_selection":
		randomIndex := rand.Intn(len(lb.config.BackendServers))
		server = lb.config.BackendServers[randomIndex]
	default:
		server = lb.config.BackendServers[0] // Fallback to the first server
	}

	server.Load++ // Simulating load increment
	return server
}

// Check if load exceeds the threshold and add a new backend server if necessary
func (lb *LoadBalancer) addServer() {
	for _, server := range lb.config.BackendServers {
		if server.Load > 5 {
			newServerID := len(lb.config.BackendServers) + 1
			newServer := &BackendServer{ID: newServerID, Load: 0, Port: fmt.Sprintf("800%d", newServerID)}
			lb.config.BackendServers = append(lb.config.BackendServers, newServer)
			log.Printf("Added new backend server %d on port %s\n", newServerID, newServer.Port)
		}
	}
}

// Health check handler to return the load of all backend servers
func (lb *LoadBalancer) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "{")
	for _, server := range lb.config.BackendServers {
		fmt.Fprintf(w, "\"server_%d\": %d,\n", server.ID, server.Load)
	}
	fmt.Fprintln(w, "}")
}

func (lb *LoadBalancer) httpHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == lb.config.HealthEndpoint {
		lb.healthCheckHandler(w, r)
		return
	}

	// next backend server
	server := lb.nextRequest()
	serverAddr := fmt.Sprintf("http://localhost:%s/b_server%d", server.Port, server.ID)

	// forward request to backend server
	resp, err := http.Get(serverAddr)
	if err != nil {
		http.Error(w, "Failed to reach backend server", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

// Backend server handler
func backendServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from Backend Server %s\n", r.URL.Path)
}

func main() {
	lb := newLoadBalancer()

	// start backend servers
	for _, server := range lb.config.BackendServers {
		go func(server *BackendServer) {
			http.HandleFunc(fmt.Sprintf("/b_server%d", server.ID), backendServerHandler)
			log.Printf("Starting backend server %d on port %s\n", server.ID, server.Port)
			log.Fatal(http.ListenAndServe(":"+server.Port, nil))
		}(server)
	}

	go lb.addServer()

	http.HandleFunc("/", lb.httpHandler)

	fmt.Printf("Starting load balancer on port %s\n", lb.port)
	log.Fatal(http.ListenAndServe(":"+lb.port, nil))
}
