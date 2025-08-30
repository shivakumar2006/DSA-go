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

func (l *DLL) InsertLast(value int) {
	newNode := &Node{data: value}
	last := l.head

	newNode.next = nil

	if l.head == nil {
		newNode.prev = nil
		l.head = newNode
		return
	}

	for last.next != nil {
		last = last.next
	}

	last.next = newNode
	newNode.prev = last
}

func (l *DLL) InsertAt(value int, index int) {
	newNode := &Node{data: value}

	// if inserting at the head
	if index == 0 {
		newNode.next = l.head
		newNode.prev = nil
		if l.head != nil {
			l.head.prev = newNode
		}
		l.head = newNode
		return
	}

	// traverse to the node just before index
	temp := l.head
	for i := 0; i < index-1 && temp != nil; i++ {
		temp = temp.next
	}

	if temp == nil {
		fmt.Println("out of range")
	}

	newNode.next = temp.next
	newNode.prev = temp

	if temp.next != nil {
		temp.next.prev = newNode
	}

	temp.next = newNode
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
