// Insert nodes in a singly linkedlist and also insertLast node in this list

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode // because, the first node is both head and tail...
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	l.tail = newNode // update newNode and make it tail
}

func (l *LinkedList) InsertLast(value int) {
	newNode := &Node{data: value}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) InsertAtIndex(value int, index int) {
	newNode := &Node{data: value}

	// if inserting at the head
	if index == 0 {
		newNode.next = l.head
		l.head = newNode
		if l.tail == nil {
			l.tail = newNode
		}
		return
	}

	// traverse to the node
	temp := l.head
	for i := 1; i < index; i++ {
		temp = temp.next
	}

	// if index is out of bound
	if temp == nil {
		fmt.Println("Index is out of range")
		return
	}

	// Insert the newNode
	newNode.next = temp.next
	temp.next = newNode

	// update tail if inserted at the end
	if newNode.next == nil {
		l.tail = newNode
	}
}

func (l *LinkedList) DeleteFirst() {
	if l.head == nil {
		fmt.Println("List is empty, nothing to delete")
		return
	}

	l.head = l.head.next

	// if the list become empty update the tail as well
	if l.head == nil {
		l.tail = nil
	}
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("End")
}

// func (l *LinkedList) Insert(value int) {
// 	newNode := &Node{data: value}
// 	if l.head == nil { // if list is empty
// 		l.head = newNode
// 		l.tail = newNode // because, first node is also a tail node
// 		return
// 	}

// 	temp := l.head
// 	for temp.next != nil {
// 		temp = temp.next
// 	}
// 	temp.next = newNode
// 	l.tail = newNode // update tail by newNode
// }

// func (l *LinkedList) InsertLast(value int) {
// 	newNode := &Node{data: value}
// 	for l.tail == nil { // if list is empty
// 		l.head = newNode
// 		l.tail = newNode
// 		return
// 	}

// 	l.tail.next = newNode
// 	l.tail = newNode
// }

// func (l *LinkedList) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("End")
// }

// package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head *Node
// 	tail *Node
// }

// // insert head
// func (l *LinkedList) Insert(value int) {
// 	newNode := &Node{data: value}
// 	if l.head == nil {
// 		l.head = newNode
// 		l.tail = newNode // because, first node is both head and tail
// 	} else {
// 		newNode.next = l.head
// 		l.head = newNode
// 	}
// }

// func (l *LinkedList) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("End")
// }

// package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head *Node
// 	tail *Node
// }

// // insert
// func (l *LinkedList) Insert(value int) {
// 	newNode := &Node{data: value}
// 	if l.head == nil {
// 		l.head = newNode
// 		return
// 	}

// 	temp := l.head
// 	for temp.next != nil {
// 		temp = temp.next
// 	}
// 	temp.next = newNode
// }

// // insert last
// func (l *LinkedList) InsertLast(value int) {
// 	newNode := &Node{data: value}
// 	if l.tail == nil {
// 		l.tail = newNode
// 		return
// 	}

// 	temp := l.tail
// 	for temp.next != nil {
// 		temp = temp.next
// 	}
// 	temp.next = newNode
// }

// // display
// func (l *LinkedList) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("nil")
// }

// // Node represents a single node
// type Node struct {
// 	data int
// 	next *Node
// }

// // Linkedlist represents a list
// type LinkedList struct {
// 	head *Node
// }

// // Inset add a new node at the end
// func (l *LinkedList) Insert(value int) {
// 	newNode := &Node{data: value}
// 	if l.head == nil {
// 		l.head = newNode
// 		return
// 	}

// 	temp := l.head
// 	for temp.next != nil {
// 		temp = temp.next
// 	}
// 	temp.next = newNode
// }

// // Display and print all new nodes
// func (l *LinkedList) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("nil")
// }
