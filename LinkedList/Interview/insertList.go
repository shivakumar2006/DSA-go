// insertion in linkedlist using recursion

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAt(value int, index int) {
	l.head = insertRec(l.head, value, index)
}

func insertRec(temp *Node, value, index int) *Node {
	if index == 0 {
		newNode := &Node{data: value}
		newNode.next = temp
		return newNode
	}

	if temp == nil {
		fmt.Println("index is out of bound")
		return nil
	}

	temp.next = insertRec(temp.next, value, index-1)
	return temp
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("end")
}

func main() {
	list := &LinkedList{}

	list.InsertAt(50, 0)
	list.InsertAt(40, 1)
	list.InsertAt(10, 2)
	list.InsertAt(60, 3)
	list.InsertAt(5, 2) // insert at the

	list.Display()
}

// func (l *LinkedList) InsertAt(value, index int) {
// 	l.head = insertRec(l.head, value, index)
// }

// func insertRec(temp *Node, value, index int) *Node {
// 	if index == 0 {
// 		newNode := &Node{data: value}
// 		newNode.next = temp
// 		return newNode
// 	}

// 	if temp == nil {
// 		fmt.Println("index is out of bound")
// 		return nil
// 	}

// 	temp.next = insertRec(temp.next, value, index-1)
// 	return temp
// }

// func (l *LinkedList) InsertAt(value int, index int) {
// 	l.head = insertRec(l.head, value, index)
// }

// func insertRec(temp *Node, value, index int) *Node {
// 	if index == 0 {
// 		newNode := &Node{data: value}
// 		newNode.next = temp
// 		return newNode
// 	}

// 	if temp == nil {
// 		fmt.Println("index is out of bound")
// 		return nil
// 	}

// 	// now move to the next node with index-1
// 	temp.next = insertRec(temp.next, value, index-1)
// 	return temp
// }

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head *Node
// }

// func (l *LinkedList) InsertAt(value int, index int) {
// 	l.head = insertRec(l.head, value, index)
// }

// func insertRec(temp *Node, value int, index int) *Node {
// 	if index == 0 {
// 		// create new node and link
// 		newNode := &Node{data: value}
// 		newNode.next = temp
// 		return newNode
// 	}

// 	if temp == nil {
// 		fmt.Println("index out of bound")
// 		return nil
// 	}

// 	// move to the next node with index-1
// 	temp.next = insertRec(temp.next, value, index-1)
// 	return temp
// }

// // function to print list
// func (l *LinkedList) Print() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("end")
// }
