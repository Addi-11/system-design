package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var MAX_INT = 100000000
var totalPrimeNumbers int32 = 0
var CONCURRENCY = 10
var currentNum int32 = 2

func checkPrime(x int) {
	if x&1 == 0 {
		return 
	}
	for i:=3; i<= int(math.Sqrt(float64(x))); i++{
		if x%i == 0{
			return
		}
	}

	atomic.AddInt32(&totalPrimeNumbers, 1)
}

func doBatch(name int, nstart int, nend int, wg *sync.WaitGroup){
	defer wg.Done()
	start := time.Now()
	for i:=nstart; i<nend; i++ {
		checkPrime(i)
	}

	fmt.Printf("Thread %s [%d, %d) completed in %s\n", strconv.Itoa(name), nstart, nend, time.Since(start))
}

func doWork(name int, wg *sync.WaitGroup){
	defer wg.Done()

	start := time.Now()

	for{
		x := atomic.AddInt32(&currentNum, 1)
		if x > int32(MAX_INT){
			break
		}
		checkPrime(int(x))
	}

	fmt.Printf("Thread %s completed in %s\n", strconv.Itoa(name), time.Since(start))
}

var wg sync.WaitGroup

func main() {
	start := time.Now()

	// APPROACH 1: Naive Compute Prime
	for i:=3; i<=MAX_INT; i++{
		checkPrime(i)
	}
	fmt.Println("APPROACH 1: Naive Implementation\nchecking till", MAX_INT, "found", totalPrimeNumbers+1, "prime numbers. Took", time.Since(start))

	// APPROACH 2: unfair thread - using batching
	start = time.Now()
	totalPrimeNumbers = 0

	nstart := 3
	batchSize := int(float64(MAX_INT)/float64(CONCURRENCY))

	for i:=0; i< CONCURRENCY; i++ {
		wg.Add(1)
		go doBatch(i, nstart, nstart+batchSize, &wg)
		nstart += batchSize
	}

	wg.Wait()
	fmt.Println("APPROACH 2: Unfair Threading\nchecking till", MAX_INT, "found", totalPrimeNumbers+1, "prime numbers. Took", time.Since(start))

	// APPROACH 3: fair thread, any free thread checks if the number is prime
	start = time.Now()
	totalPrimeNumbers = 0
	
	for i:=0; i< CONCURRENCY; i++ {
		wg.Add(1)
		go doWork(i, &wg)
	}
	
	wg.Wait()
	fmt.Println("APPROACH 3: Fair Threading\nchecking till", MAX_INT, "found", totalPrimeNumbers+1, "prime numbers. Took", time.Since(start))
}
