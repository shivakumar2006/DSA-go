// // Reorder list
// // Google, Facebook question

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

func middleNode(head *Node) *Node {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func reverseList(head *Node) *Node {
	var prev *Node
	current := head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	return prev
}

func reorderList(head *Node) {
	if head == nil || head.next == nil {
		return
	}

	mid := middleNode(head)
	headSecond := reverseList(mid.next)
	mid.next = nil // break into half

	headFirst := head

	for headSecond != nil {
		temp1 := headFirst.next
		temp2 := headSecond.next

		headFirst.next = headSecond
		headSecond.next = temp1

		headFirst = temp1
		headSecond = temp2
	}
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("END")
}

func main() {
	list := &LinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	list.Insert(5)

	fmt.Print("Original list : ")
	list.Display()

	reorderList(list.head)

	fmt.Print("reordered list : ")
	list.Display()
}

// package main

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head *Node
// }

// func (l *LinkedList) Insert(value int) {
// 	newNode := &Node{data: value}

// 	if l.head == nil {
// 		l.head = newNode
// 		return
// 	}

// 	temp := l.head
// 	for temp.next != nil {
// 		temp = temp.next
// 	}
// 	temp.next = newNode
// }

// func middleNode(head *Node) *Node {
// 	slow := head
// 	fast := head

// 	for fast != nil && fast.next != nil {
// 		slow = slow.next
// 		fast = fast.next.next
// 	}
// 	return slow
// }

// func reverseList(head *Node) *Node {
// 	var prev *Node
// 	current := head

// 	for current != nil {
// 		next := current.next
// 		current.next = prev
// 		prev = current
// 		current = next
// 	}
// 	return prev
// }

// func reorderList(head *Node) {
// 	if head == nil || head.next == nil {
// 		return
// 	}

// 	mid := middleNode(head)
// 	headSecond := reverseList(mid.next)
// 	mid.next = nil // break first half

// 	headFirst := head

// 	for headSecond != nil {
// 		temp1 := headFirst.next
// 		temp2 := headSecond.next

// 		headFirst.next = headSecond
// 		headSecond.next = temp1

// 		headFirst = temp1
// 		headSecond = temp2
// 	}
// }

// func (l *LinkedList) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("END")
// }

// func main() {
// 	list := &LinkedList{}

// 	list.Insert(1)
// 	list.Insert(2)
// 	list.Insert(3)
// 	list.Insert(4)
// 	list.Insert(5)

// 	fmt.Print("original list : ")
// 	list.Display()

// 	reorderList(list.head)

// 	fmt.Print("reorder list : ")
// 	list.Display()
// }

// // func reorderList(head *Node) {
// // 	if head == nil || head.next == nil {
// // 		return
// // 	}

// // 	// 1. Find middle
// // 	mid := middleNode(head)
// // 	headSecond := reverseList(mid.next)
// // 	mid.next = nil // break first half

// // 	headFirst := head

// // 	// 2. Merge two halves
// // 	for headSecond != nil {
// // 		temp1 := headFirst.next
// // 		temp2 := headSecond.next

// // 		headFirst.next = headSecond
// // 		headSecond.next = temp1

// // 		headFirst = temp1
// // 		headSecond = temp2
// // 	}
// // }
