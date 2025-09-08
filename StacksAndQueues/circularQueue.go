// // Circular Queue

package main

import (
	"errors"
	"fmt"
)

type CircularQueue struct {
	data  []int
	size  int
	front int
	end   int
	count int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, capacity),
		size:  capacity,
		front: 0,
		end:   0,
		count: 0,
	}
}

// check if the queue is full or not
func (q *CircularQueue) isFull() bool {
	return q.count == q.size
}

// check if the queue is empty or not
func (q *CircularQueue) isEmpty() bool {
	return q.count == 0
}

// now insert
func (q *CircularQueue) Insert(value int) bool {
	if q.isFull() {
		return false
	}
	q.data[q.end] = value
	q.end = (q.end + 1) % q.size
	q.count++
	return true
}

// remove (dequeue)
func (q *CircularQueue) Remove() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty.")
	}
	removed := q.data[q.front]
	q.front = (q.front + 1) % q.size
	q.count--
	return removed, nil
}

// display the queue
func (q *CircularQueue) Display() {
	if q.isEmpty() {
		fmt.Println("Queue is Empty!")
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

	value, _ := queue.Remove()
	fmt.Println("Removed : ", value)
	queue.Display()
}

// package main

// import "fmt"

// type CircularQueue struct {
// 	data  []int
// 	size  int
// 	front int
// 	end   int
// 	count int
// }

// // Constructor with default size
// func NewCircularQueue(capacity int) *CircularQueue {
// 	return &CircularQueue{
// 		data:  make([]int, capacity),
// 		size:  capacity,
// 		front: 0,
// 		end:   0,
// 		count: 0,
// 	}
// }

// // check if queue is full or not
// func (q *CircularQueue) isFull() bool {
// 	return q.count == q.size
// }

// // check if the queue is empty or not
// func (q *CircularQueue) isEmpty() bool {
// 	return q.count == 0
// }

// // Insert (enqueue)
// func (q *CircularQueue) Insert(value int) bool {
// 	if q.isFull() {
// 		return false
// 	}
// 	q.data[q.end] = value
// 	q.end = (q.end + 1) % q.size
// 	q.count++
// 	return true
// }

// func (q *CircularQueue) Display() {
// 	if q.isEmpty() {
// 		fmt.Println("Empty")
// 		return
// 	}
// 	i := q.front
// 	for {
// 		fmt.Printf("%d -> ", q.data[i])
// 		i = (i + 1) % q.size
// 		if i == q.end {
// 			break
// 		}
// 	}
// 	fmt.Println("END")
// }

// func main() {
// 	queue := NewCircularQueue(5)

// 	queue.Insert(10)
// 	queue.Insert(20)
// 	queue.Insert(30)
// 	queue.Insert(40)
// 	queue.Insert(50)

// 	queue.Display()
// }
