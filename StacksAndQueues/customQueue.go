// Custom queue

package main

import (
	"errors"
	"fmt"
)

type CustomQueue struct {
	data []int
	size int
}

// create new queues
func NewCustomQueue(capacity int) *CustomQueue {
	return &CustomQueue{
		data: make([]int, 0, capacity),
		size: capacity,
	}
}

// Enqueue insert at end
func (cq *CustomQueue) Enqueue(value int) error {
	if len(cq.data) >= cq.size {
		return errors.New("Queue is full!")
	}
	cq.data = append(cq.data, value)
	return nil
}

func main() {
	cq := NewCustomQueue(3)

	// Insert elements
	_ = cq.Enqueue(10)
	_ = cq.Enqueue(20)
	_ = cq.Enqueue(30)

	fmt.Println("Queue : ", cq.data)
}
