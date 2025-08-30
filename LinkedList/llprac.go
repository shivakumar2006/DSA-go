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

func (l *LinkedList) InsertAll(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
	l.tail = newNode // make newNode atail
}

func (l *LinkedList) InsertEnd(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) InsertAt(value int, index int) {
	newNode := &Node{data: value}

	// if inserting at the end
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
	for i := 0; i < index-1; i++ {
		temp = temp.next
	}

	// check index is out of bound
	if temp == nil {
		fmt.Println("Out of bound")
	}

	// insert the new node
	newNode.next = temp.next
	temp.next = newNode

	// if insert at the end then
	if newNode.next == nil {
		l.tail = newNode
	}
}

func (l *LinkedList) DisplayAll() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("End")
}

func main() {
	list := LinkedList{}

	list.InsertAll(10)
	list.InsertAll(8)
	list.InsertAll(20)
	list.InsertAll(16)
	list.DisplayAll()

	fmt.Println("before insert tail node...")
	list.InsertEnd(100)
	fmt.Println("After insert tail node...")
	list.DisplayAll()
	list.InsertAt(120, 2)
	list.DisplayAll()
}
