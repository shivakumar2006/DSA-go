// // Middle of the linkedlist
// // Google question

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

func (l *LinkedList) findMiddle() *Node {
	if l.head == nil {
		return nil
	}

	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	return slow
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

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)

	list.Display()

	middle := list.findMiddle()
	if middle != nil {
		fmt.Println("the middle element of the list is : ", middle.data)
	} else {
		fmt.Println("list is empty")
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
// }

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

// func (l *LinkedList) findMiddle() *Node {
// 	if l.head == nil {
// 		return nil
// 	}

// 	fast := l.head
// 	slow := l.head

// 	for fast != nil && fast.next != nil {
// 		fast = fast.next.next
// 		slow = slow.next
// 	}

// 	return slow
// }

// func (l *LinkedList) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("End")
// }

// func main() {
// 	list := &LinkedList{}

// 	list.Insert(1)
// 	list.Insert(2)
// 	list.Insert(3)
// 	list.Insert(4)
// 	list.Insert(5)

// 	list.Display()

// 	middle := list.findMiddle()
// 	if middle != nil {
// 		fmt.Println("middle node value : ", middle.data)
// 	} else {
// 		fmt.Println("list is empty")
// 	}
// }
