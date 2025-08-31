// Remove duplicafes from sorted linkedlist

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
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) Duplicates() {
	temp := l.head

	for temp.next != nil {
		if temp.data == temp.next.data {
			temp.next = temp.next.next
		} else {
			temp = temp.next
		}
	}

	l.tail = temp
	l.tail.next = nil
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
	list.Insert(1)
	list.Insert(1)
	list.Insert(2)
	list.Insert(4)
	list.Insert(4)
	list.Display()
	list.Duplicates()
	list.Display()
}
