// Insert nodes in a singly linkedlist and also insertLast node in this list

package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Inset(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode // because, first node is also a tail node
		return
	}

}

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
