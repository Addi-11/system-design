package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
	"bytes"
	"sync/atomic"
)

var (
	machineID = generateMachineID() // 3 bytes for machine ID
	counter = uint32(rand.Intn(1<<24)) // 3 bytes for counter
	counterLock = sync.Mutex{}
)

func generateMachineID() []byte{
	hostname, err := os.Hostname()
	if err != nil{
		hostname = "localhost"
	}
	h := hash([]byte(hostname))
	return h[:3]
}

func hash(data []byte) []byte{
	var result [3]byte
	for i, b := range data{
		result[i%3] ^= b
	}
	return result[:]
}

func objectIDHex(id []byte) string{
	return hex.EncodeToString(id)
}

func generateObjectID() []byte{

	buf := new(bytes.Buffer)

	timestamp := uint32(time.Now().Unix())
	pid := os.Getpid()
	counter := atomic.AddUint32(&counter, 1) & 0xFFFFFF

	buf.Write([]byte{
		byte(timestamp >> 24),
		byte(timestamp >> 16),
		byte(timestamp >> 8),
		byte(timestamp),
	})

	buf.Write(machineID)

	buf.Write([]byte{
		byte(pid >> 8),
		byte(pid),
	})

	buf.Write([]byte{
		byte(counter >> 16),
		byte(counter >> 8),
		byte(counter),
	})

	return buf.Bytes()
}

func generateMultipleIDs(numIDs int, results chan<- string, wg *sync.WaitGroup){
	defer wg.Done()

	for i := 0; i < numIDs; i++ {
		id := generateObjectID()
		results <- objectIDHex(id)
	}
}

func main(){
	fmt.Printf("Machine ID: %s\n", machineID)
	
	const numGoRoutines = 5
	const ids = 3
	totalIDs := numGoRoutines * ids

	results := make(chan string, totalIDs)
	var wg sync.WaitGroup

	for i := 0; i < numGoRoutines; i++ {
		wg.Add(1)
		go generateMultipleIDs(ids, results, &wg)
	}

	wg.Wait()
	close(results)

	for id := range results {
		fmt.Println(id)
	}
}