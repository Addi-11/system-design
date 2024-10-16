package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup
var count int = 0

func incCount(){
	// we want only one thread to excute it at a time
	mu.Lock()
	count++;
	mu.Unlock()
	wg.Done()
}

func doCount(){
	for i:=0; i < 1000000; i++{
		wg.Add(1)
		go incCount()
	}
}

func main(){
	count = 0
	doCount()
	// wait for all the threads to complete before printing the result
	wg.Wait()
	fmt.Println(count)
}