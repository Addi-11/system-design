package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	machineID string
	filename string = "counter.txt"
	mu sync.Mutex
	wg sync.WaitGroup
	batchSize = 500
)

func saveCounter(counter int){
	data := strconv.Itoa(counter)
	os.WriteFile(filename, []byte(data), 0644)
}

func loadCounter() int{
	data, err := os.ReadFile(filename)
	if err != nil{
		panic(err)
	}

	id, err := strconv.Atoi(string(data))
	if err != nil{
		panic(err)
	}
	return id
}


// returns batch id's in sets of 500
func getBatchID(prefix string) []string{
	ids := make([]string, batchSize)
	counterBase := loadCounter()
	wg.Add(batchSize)

	for i:=0; i<batchSize; i++{
		go func(i int){
			defer wg.Done()
			id := fmt.Sprintf("%s-%s-%d-%d", prefix, machineID, counterBase+i, time.Now().UnixNano())
			ids[i] = id
		}(i)
	}

	wg.Wait()
	saveCounter(counterBase + batchSize)
	return ids
}

// generate random machine ID for each request - simulate multiple
func generateMachineID() string {
	return fmt.Sprintf("M%d", rand.Intn(100)+1)
}

func orderHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "NEW ORDER IDS")
	ids := getBatchID("orders")
	for _, id := range(ids){
		fmt.Fprintln(w, id)
	}
}

func paymentHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "NEW PAYMENT IDS")
	ids := getBatchID("payments")
	for _, id := range(ids){
		fmt.Fprintln(w, id)
	}
}

func main(){
	machineID = generateMachineID()
	
	http.HandleFunc("/orders/get_ids", orderHandler)
	http.HandleFunc("/payments/get_ids", paymentHandler)

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}