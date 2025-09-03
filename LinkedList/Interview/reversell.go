// reverse the linked list using recursion

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

	list.Insert(5)
	list.Insert(3)
	list.Insert(1)
	list.Insert(2)
	list.Insert(4)

	fmt.Print("Original list : ")
	list.Display()
}
