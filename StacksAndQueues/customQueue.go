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

// Dequeue (remove from front)
func (cq *CustomQueue) Dequeue() (int, error) {
	if len(cq.data) == 0 {
		return 0, errors.New("Queue is empty!")
	}
	value := cq.data[0]
	cq.data = cq.data[1:] // remove first element

	return value, nil
}

func (cq *CustomQueue) Peek() (int, error) {
	if len(cq.data) == 0 {
		return 0, errors.New("Queue is empty")
	}
	return cq.data[0], nil
}

func main() {
	cq := NewCustomQueue(3)

	// Insert elements
	_ = cq.Enqueue(10)
	_ = cq.Enqueue(20)
	_ = cq.Enqueue(30)

	fmt.Println("Queue : ", cq.data)

	// try to insert more than capacity
	err := cq.Enqueue(40)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	// remove elements
	val, _ := cq.Dequeue()
	fmt.Println("Dequed : ", val)
	fmt.Println("Queue after dequeue : ", cq.data)

	// Peek from element
	front, _ := cq.Peek()
	fmt.Println("Front Element : ", front)
}
