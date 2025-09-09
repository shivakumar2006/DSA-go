// // implement queues using stacks
// // Google interview question

package main

import (
	"fmt"
)

type queueUsingStack struct {
	stack1 []int
	stack2 []int
}

func NewQueueUsingStack() *queueUsingStack {
	return &queueUsingStack{
		stack1: []int{},
		stack2: []int{},
	}
}

// first add or Enqueue all the elements in the stack 1
func (q *queueUsingStack) Enqueue(value int) {
	q.stack1 = append(q.stack1, value)
}

// check if its empty
func (q *queueUsingStack) isEmpty() bool {
	return len(q.stack1) == 0 && len(q.stack2) == 0
}

// now display
func (q *queueUsingStack) Display() {
	if q.isEmpty() {
		fmt.Println("Empty!")
		return
	}

	// print stack 2
	for i := len(q.stack2) - 1; i >= 0; i-- {
		fmt.Printf("%d -> ", q.stack2[i])
	}

	// print stack 1
	for i := 0; i < len(q.stack1); i++ {
		fmt.Printf("%d -> ", q.stack1[i])
	}

	fmt.Println("END")
}

func main() {
	q := NewQueueUsingStack()

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	q.Display()
}

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// // Queue using stack implements a queue using two stacks
// type queueUsingStack struct {
// 	stack1 []int // main stack
// 	stack2 []int // helper stack
// }

// // creates a new empty queue
// func NewQueueUsingStack() *queueUsingStack {
// 	return &queueUsingStack{
// 		stack1: []int{},
// 		stack2: []int{},
// 	}
// }

// // Enqueue adds an element using queue
// func (q *queueUsingStack) Enqueue(value int) {
// 	q.stack1 = append(q.stack1, value)
// }

// // check if the queue is empty
// func (q *queueUsingStack) isEmpty() bool {
// 	return len(q.stack1) == 0 && len(q.stack2) == 0
// }

// // Dequeue remove the front element, means first element from stack 1 which is 30 then 20 then 10
// func (q *queueUsingStack) Dequeue() (int, error) {
// 	if q.isEmpty() {
// 		return 0, errors.New("Queue is Empty!")
// 	}

// 	// if stack 2 is empty move all the elements to stack 2 from stack 1.
// 	if len(q.stack2) == 0 {
// 		for len(q.stack1) > 0 {
// 			n := len(q.stack1)
// 			top := q.stack1[n-1]
// 			q.stack1 = q.stack1[:n-1]
// 			q.stack2 = append(q.stack2, top)
// 		}
// 	}

// 	// pop from stack 2
// 	n := len(q.stack2)
// 	front := q.stack2[n-1]
// 	q.stack2 = q.stack2[:n-1]
// 	return front, nil
// }

// // retunrs the front element without removing it.
// func (q *queueUsingStack) Front() (int, error) {
// 	if q.isEmpty() {
// 		return 0, errors.New("queue is empty")
// 	}

// 	if len(q.stack2) == 0 {
// 		for len(q.stack1) > 0 {
// 			n := len(q.stack1)
// 			top := q.stack1[n-1]
// 			q.stack1 = q.stack1[:n-1]
// 			q.stack2 = append(q.stack2, top)
// 		}
// 	}
// 	return q.stack2[len(q.stack2)-1], nil
// }

// // Display prints the element of the queue
// func (q *queueUsingStack) Display() {
// 	if q.isEmpty() {
// 		fmt.Println("Empty!")
// 		return
// 	}

// 	// print stack 2 in reverse
// 	for i := len(q.stack2) - 1; i >= 0; i-- {
// 		fmt.Printf("%d -> ", q.stack2[i])
// 	}

// 	// print stack 1 in normal order
// 	for i := 0; i < len(q.stack1); i++ {
// 		fmt.Printf("%d -> ", q.stack1[i])
// 	}

// 	fmt.Println("END")
// }

// func main() {
// 	q := NewQueueUsingStack()

// 	q.Enqueue(10)
// 	q.Enqueue(20)
// 	q.Enqueue(30)
// 	q.Display() // 10 -> 20 -> 30 -> END

// 	value, _ := q.Dequeue()
// 	fmt.Println("Dequed : ", value)

// 	q.Display() // 20 -> 30 -> END

// 	front, _ := q.Front()
// 	fmt.Println("Front : ", front)

// }
