package main

import (
	"math/rand"
	"log"
	"fmt"
	"sync"
	"time"
)

type BlockingQueue struct {
	queue []int32
	mu sync.Mutex
	notEmpty *sync.Cond // condition var for Dequeue to wait, when queue is empty
	notFull *sync.Cond // condition var for Enqueue to wait, when queue is full
	capacity int32

}

func NewBlockingQueue (capacity int32) *BlockingQueue{
	q := &BlockingQueue{
		queue : make([]int32, 0, capacity),
		capacity: capacity,
	}

	// condition relies on the mutex, for thread safe synchronoisation
	q.notEmpty = sync.NewCond(&q.mu)
	q.notFull = sync.NewCond(&q.mu)
	return q
}

func (q *BlockingQueue) Enqueue(item int32){
	q.mu.Lock()
	defer q.mu.Unlock()

	// wait until there is space in the queue
	for len(q.queue) >= int(q.capacity) {
		q.notFull.Wait()
	}

	q.queue = append(q.queue, item)

	// signal waiting deque routine that queue is no longer empty
	q.notEmpty.Signal()

	log.Println("Enqueued:", item)
}

func (q *BlockingQueue) Dequeue() int32 {
	q.mu.Lock()
	defer q.mu.Unlock()

	// wait till there is something in the queue
	for len(q.queue) == 0 {
		q.notEmpty.Wait()
	}

	item := q.queue[0]
	q.queue = q.queue[1:]

	// signal enqueue operation, queue is no longer full
	q.notFull.Signal()

	log.Println("Dequeued:", item)

	return item
}

func (q *BlockingQueue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.queue)
}

var wgE sync.WaitGroup
var wgD sync.WaitGroup

func main(){
	q := NewBlockingQueue(50)

	for i := 0; i< 100; i++{
		wgE.Add(1)
		go func(){
			q.Enqueue(rand.Int31())
			wgE.Done()
		}()
	}

	for i:=0; i<100; i++{
		wgD.Add(1)
		go func (){
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
			q.Dequeue()
			wgD.Done()
		}()
	}

	wgE.Wait()
	wgD.Wait()

	fmt.Println("\n\nQueue size: ", q.Size())
}