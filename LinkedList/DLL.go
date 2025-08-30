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
		l.head.prev = newNode
	}
	l.head = newNode
}

func (l *DLL) Display() {
	temp := l.head
	var last *Node
	for temp != nil {
		fmt.Print(temp.data, " <-> ")
		last = temp
		temp = temp.next
	}
	fmt.Println("end")

	fmt.Println("print in reverse")
	for last != nil {
		fmt.Print(last.data, " <-> ")
		last = last.prev
	}
	fmt.Println("start")
}
