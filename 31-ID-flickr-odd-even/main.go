package main

import (
	"fmt"
	"sync"
	"time"
)

type IDGenerator struct{
	even int64
	odd int64
	isOdd bool
	mu sync.Mutex
}

func NewIDGenerator(even int64, odd int64) *IDGenerator{
	return &IDGenerator{
		even: even,
		odd: odd,
		isOdd: true,
	}
}

func (gen *IDGenerator) GenerateID() string{
	var id int64
	var prefix string
	gen.mu.Lock()
	defer gen.mu.Unlock()

	time.Sleep(10 * time.Millisecond)

	if gen.isOdd {
		id = gen.even
		prefix = "even"
		gen.even += 2
	} else {
		id = gen.odd
		prefix = "odd"
		gen.odd += 2
	}
	gen.isOdd = !gen.isOdd
	return fmt.Sprintf("%s-%d", prefix, id)
}

func main(){
	idGen := NewIDGenerator(100, 101)
	var wg sync.WaitGroup

	numGoRoutine := 10
	numIds := 7
	idMap := sync.Map{}

	// Goroutines = servers making parallel request for ID generation
	for i := 0; i < numGoRoutine; i++ {
		wg.Add(1)
		go func(routineID int) {
			defer wg.Done()
			for j := 0; j < numIds; j++ {
				id := idGen.GenerateID()
				idMap.Store(id, routineID)
				fmt.Printf("Routine: %d, Generated ID: %s\n", routineID, id)
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("\nVerifying generated IDs...")
	var allIDs []string
	idMap.Range(func(key, value any) bool {
		allIDs = append(allIDs, key.(string))
		return true
	})

	unique := make(map[string]bool)
	for _, id := range allIDs {
		if unique[id] {
			fmt.Printf("Duplicate ID found: %s\n", id)
		} else {
			unique[id] = true
		}
	}
	fmt.Printf("Total IDs generated: %d\n", len(allIDs))
	fmt.Printf("All IDs are unique: %t\n", len(allIDs) == len(unique))
}