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

// Detect if cycle is present
func (l *LinkedList) hasCycle() bool {
	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			return true
		}
	}
	return false
}

// Calculate cycle length starting from meeting node
func (l *LinkedList) cycleLength(node *Node) int {
	temp := node.next
	length := 1
	for temp != node {
		temp = temp.next
		length++
	}
	return length
}

// Find the start node of the cycle
func (l *LinkedList) detectCycle() *Node {
	fast := l.head
	slow := l.head
	length := 0

	// Detect cycle
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			length = l.cycleLength(slow)
			break
		}
	}

	if length == 0 {
		return nil // No cycle
	}

	// Move one pointer ahead by 'length'
	f := l.head
	s := l.head
	for i := 0; i < length; i++ {
		s = s.next
	}

	// Move both one step at a time until they meet
	for f != s {
		f = f.next
		s = s.next
	}

	return s
}

// Print linked list (will loop infinitely if cycle exists)
func (l *LinkedList) Display(limit int) {
	temp := l.head
	count := 0
	for temp != nil && count < limit { // limit to avoid infinite loop
		fmt.Print(temp.data, " -> ")
		temp = temp.next
		count++
	}
	fmt.Println("...")
}

func main() {
	list := &LinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)
	list.Insert(6)

	// create a cycle for testing (cycle starting at node 3)
	temp := list.head
	var thirdNode *Node
	for temp != nil {
		if temp.data == 3 {
			thirdNode = temp
		}
		if temp.next == nil {
			temp.next = thirdNode // create the cycle
			break
		}
		temp = temp.next
	}

	if list.hasCycle() {
		fmt.Println("Cycle is present in the list")
		length := list.cycleLength(thirdNode)
		fmt.Println("Cycle length:", length)

		startNode := list.detectCycle()
		if startNode != nil {
			fmt.Println("Cycle starts at node with value:", startNode.data)
		}
	} else {
		fmt.Println("No cycle detected")
	}

	// Display list safely with limit
	list.Display(15)
}
