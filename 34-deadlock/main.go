package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type Record struct{
	mu sync.Mutex
	data string
}

var DB []Record

func aquireLock(txn string, recordIdx int){
	fmt.Printf("Transaction %s wants to aqcuire lock on Record %d\n", txn, recordIdx)
	DB[recordIdx].mu.Lock()
	fmt.Printf("Transaction %s aqcuired lock on Record %d\n", txn, recordIdx)
}

func releaseLock(txn string, recordIdx int){
	DB[recordIdx].mu.Unlock()
	fmt.Printf("Transaction %s released lock on Record %d\n", txn, recordIdx)
}

func mimicLoad(txn string, wg *sync.WaitGroup){
	defer wg.Done()

	for {
		row1 := rand.Intn(len(DB))
		row2 := rand.Intn(len(DB))

		if row1 == row2{
			continue
		}

		if row1 > row2{
			tmp := row1
			row1 = row2
			row2 = tmp
		}

		// lock rows
		aquireLock(txn, row1)
		aquireLock(txn, row2)

		time.Sleep(2 * time.Second)

		// unlock rows
		releaseLock(txn, row1)
		releaseLock(txn, row2)
	}
}

func main(){
	numTransactions := 3
	numRecords := 5
	DB = make([]Record, numRecords)
	
	var wg sync.WaitGroup

	// initialize DB
	for i:=1; i<=numRecords; i++{
		DB = append(DB, Record{data: fmt.Sprintf("record_%f", i)})
	}

	for i:=1; i<=numTransactions; i++{
		wg.Add(1)
		go mimicLoad(fmt.Sprintf("T%d", i), &wg)
	}

	wg.Wait()
	fmt.Printf("All transactions complete.\n")
}
