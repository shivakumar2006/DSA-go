// Custom stack

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

// push adds an element in the top of the stack.
func (s *Stack) Push(value int) {
	newNode := &Node{data: value}
	newNode.next = s.top
	s.top = newNode
}

// pop remove the top element from the stack.
func (s *Stack) Pop() int {
	if s.top == nil {
		fmt.Println("Stack is empty!")
		return -1
	}
	value := s.top.data
	s.top = s.top.next // just remove the top element in the stack and make head next value so now head is 40 and 50 will be removed
	return value
}

// Peek return. the top element without removing it
func (s *Stack) Peek() int {
	if s.top == nil {
		fmt.Println("Stack is empty!")
		return -1
	}
	return s.top.data
}

// Display and prints all the elements of the stack
func (s *Stack) Display() {
	temp := s.top
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("END")
}

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	fmt.Print("Stack : ")
	stack.Display()

	fmt.Println("Popped : ", stack.Pop())
	fmt.Print("Stack after pop : ")
	stack.Display()

	fmt.Println("Top elements : ", stack.Peek())
}
