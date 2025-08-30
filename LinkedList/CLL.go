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

func (l *CLL) Delete(index int) {
	if l.head == nil {
		fmt.Println("list is empty")
	}

	// delete first node
	if index == 0 {
		if l.head == l.tail {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
			l.tail.next = l.head
		}
		return
	}

	// traverse the node
	temp := l.head
	for i := 0; i < index-1; i++ {
		temp = temp.next
		if temp.next == l.head {
			fmt.Println("out of range")
			return
		}
	}

	toDelete := temp.next
	temp.next = toDelete.next

	// if deleting tail, update it
	if toDelete == l.tail {
		l.tail = temp
	}
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
