// Rotate Linkedlist

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func rotateRight(head *Node, k int) *Node {
	if k <= 0 || head == nil || head.next == nil {
		return head
	}

	last := head
	length := 1 // find length of the list first
	for last.next != nil {
		last = last.next
		length++
	}
	last.next = head

	rotations := k % length
	skip := length - rotations
	newLast := head
	for i := 0; i < skip-1; i++ {
		newLast = newLast.next
	}
	head = newLast.next
	newLast.next = nil

	return head
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("END")
}

func main() {
	list := &LinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	fmt.Print("Original list : ")
	list.Display()

	list.head = rotateRight(list.head, 2)

	fmt.Print("Rotate list : ")
	list.Display()
}

// func rotateRight(head *Node, k int) *Node {
// 	if k <= 0 || head == nil || head.next == nil {
// 		return head
// 	}

// 	last := head
// 	length := 1
// 	for last.next != nil {
// 		last = last.next
// 		length++
// 	}
// 	// because goes from end to start (5 to 1 again)
// 	last.next = head

// 	rotations := k % length    // 2 % 5 = 2
// 	skip := length - rotations // 5 - 2 = 3
// 	newLast := head            // means 1 right now
// 	for i := 0; i < skip-1; i++ {
// 		newLast = newLast.next
// 	}
// 	head = newLast.next
// 	newLast.next = nil

// 	return head
// }

// func rotateRight(head *Node, k int) *Node {
// 	if k <= 0 || head == nil || head.next == nil {
// 		return head
// 	}

// 	last := head
// 	length := 1
// 	for last.next != nil {
// 		last = last.next
// 		length++
// 	}

// 	last.next = head
// 	rotations := k % length
// 	skip := length - rotations
// 	newLast := head
// 	for i := 0; i < skip-1; i++ {
// 		newLast = newLast.next
// 	}
// 	head = newLast.next
// 	newLast.next = nil

// 	return head
// }
