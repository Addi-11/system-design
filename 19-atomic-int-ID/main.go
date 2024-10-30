package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)


func save_counter(){
	data := strconv.Itoa(counter)
	os.WriteFile(filename, []byte(data), 0644)
}

func load_counter() int{
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

func get_id(machine_id int) string{
	mu.Lock()
		counter++
	mu.Unlock()
	return fmt.Sprintf("%d-%d", machine_id, counter)
}

var (
	filename string = "counter.txt"
	counter int
	mu sync.Mutex
	wg sync.WaitGroup
)


func main(){
	counter = load_counter() + 1000 // add buffer
	start := time.Now()
	n := 100000

	for i:=0; i<n; i++{
		machine_id := rand.Intn(3) + 1 // generate a random machine id
		wg.Add(1)
		
		go func(machine_id int) {
			defer wg.Done()
			// fmt.Println("ID: ", get_id(machine_id))

			if counter%100 == 0{
				mu.Lock()
				save_counter()
				mu.Unlock()								
			}
		}(machine_id)
	}

	// save final counter
	wg.Wait()
	fmt.Printf("Time taken to generate %d IDs : %v", n, time.Since(start))
	save_counter()
}