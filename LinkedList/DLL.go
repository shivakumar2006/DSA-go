// Doubly Linkedlist

package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DLL struct {
	head *Node
}

func (l *DLL) Insert(value int) {
	newNode := &Node{data: value}

	newNode.next = l.head
	newNode.prev = nil
	if l.head != nil {
		l.head.prev = l.head
	}
	l.head = newNode
}

func (l *DLL) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " <-> ")
		temp = temp.next
	}
	fmt.Println("end")
}
