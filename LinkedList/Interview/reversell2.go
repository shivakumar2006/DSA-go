// // Reverse linked list 2
// // Google, ,microsoft, facebook question

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

func (l *LinkedList) reverseBetween(head *Node, left, right int) *Node {
	if left == right {
		return head
	}

	// iterate till left
	current := head
	var prev *Node = nil
	for i := 0; i < left-1; i++ {
		prev = current
		current = current.next
	}

	last := prev
	newEnd := current

	// now rotate between the left and right part of the list
	var next *Node = current.next
	for i := 0; current != nil && i < right-left+1; i++ {
		current.next = prev
		prev = current
		current = next
		if next != nil {
			next = next.next
		}
	}

	if last != nil {
		last.next = prev
	} else {
		head = prev
	}

	newEnd.next = current
	return head
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("End")
}

func main() {
	list := &LinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	fmt.Print("original list : ")
	list.Display()

	list.head = list.reverseBetween(list.head, 2, 4)

	fmt.Print("reverse between 2 and 4 : ")
	list.Display()
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

// func (l *LinkedList) reverseBetween(head *Node, left, right int) *Node {
// 	if left == right {
// 		return head
// 	}

// 	// skip the first left-1 node
// 	current := head
// 	var prev *Node = nil
// 	for i := 0; current != nil && i < left-1; i++ {
// 		prev = current
// 		current = current.next
// 	}

// 	last := prev
// 	newEnd := current

// 	var next *Node = current.next
// 	// reverse between left and right
// 	for i := 0; current != nil && i < right-left+1; i++ {
// 		current.next = prev
// 		prev = current
// 		current = next
// 		if next != nil {
// 			next = next.next
// 		}
// 	}

// 	// reconnect 1 with 4 and 3 with 5
// 	if last != nil {
// 		last.next = prev
// 	} else {
// 		head = prev
// 	}

// 	newEnd.next = current
// 	return head
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

// 	fmt.Print("Original list: ")
// 	list.Display()

// 	// Reverse between position 2 and 4
// 	list.head = list.reverseBetween(list.head, 2, 4)

// 	fmt.Print("After reversing between 2 and 4: ")
// 	list.Display()
// }
