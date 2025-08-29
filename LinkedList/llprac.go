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
	if temp != nil {
		temp = temp.next
	}

	l.tail.next = newNode
	l.tail = newNode // make newNode atail
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
}
