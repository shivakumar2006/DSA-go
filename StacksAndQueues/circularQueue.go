// Circular Queue

package main

import "fmt"

type CircularQueue struct {
	data  []int
	size  int
	front int
	end   int
	count int
}

// Constructor with default size
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, capacity),
		size:  capacity,
		front: 0,
		end:   0,
		count: 0,
	}
}

// check if queue is full or not
func (q *CircularQueue) isFull() bool {
	return q.count == q.size
}

// check if the queue is empty or not
func (q *CircularQueue) isEmpty() bool {
	return q.count == 0
}

// Insert (enqueue)
func (q *CircularQueue) Insert(value int) bool {
	if q.isFull() {
		return false
	}
	q.data[q.end] = value
	q.end = (q.end + 1) % q.size
	q.count++
	return true
}

func (q *CircularQueue) Display() {
	if q.isEmpty() {
		fmt.Println("Empty")
		return
	}
	i := q.front
	for {
		fmt.Printf("%d -> ", q.data[i])
		i = (i + 1) % q.size
		if i == q.end {
			break
		}
	}
	fmt.Println("END")
}

func main() {
	queue := NewCircularQueue(5)

	queue.Insert(10)
	queue.Insert(20)
	queue.Insert(30)
	queue.Insert(40)
	queue.Insert(50)

	queue.Display()
}
