package main

import (
	"fmt"
	"time"
)

type BlockingQueue struct {
	channel chan interface{}
}

func NewBlockingQueue(size int) *BlockingQueue{
	return &BlockingQueue {
		channel: make(chan interface{}, size), // buffer channel of any size
	}
}

func (q *BlockingQueue) Enqueue(item interface{}){
	q.channel <- item
}

func (q *BlockingQueue) Dequeue() interface{} {
	return <-q.channel
}

func (q *BlockingQueue) Size() int {
	return len(q.channel)
}

func main() {
	q := NewBlockingQueue(3)

	// q.Enqueue(42)
	// q.Enqueue("danger")

	// fmt.Println(q.Dequeue())
    // fmt.Println(q.Dequeue())

	go func() {
		for i := 1; i <= 5; i++ {
			q.Enqueue(i)
			fmt.Println("Enqueued:", i, " | Current size:", q.Size())
			time.Sleep(1 * time.Second) // Simulate some work
		}
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(2 * time.Second) // Simulate slower work to show blocking
			value := q.Dequeue()
			fmt.Println("Dequeued:", value, " | Current size:", q.Size())
		}

	}()

	time.Sleep(10*time.Second)
	fmt.Println(q.Size())
}