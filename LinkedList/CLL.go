package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type CLL struct {
	head *Node
	tail *Node
}

func (l *CLL) Insert(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	newNode.next = l.head

	l.tail = newNode
}

func (l *CLL) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
		if temp == l.head {
			break
		}
	}
	fmt.Println("(head)")
}
