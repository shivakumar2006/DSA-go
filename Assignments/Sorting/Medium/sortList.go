// Sort list
// https://leetcode.com/problems/sort-list/description/

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SL struct {
	head *Node
}

func (l *SL) Insert(value int) {
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

func sortList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	mid := getMid(head)
	right := mid.next
	mid.next = nil

	left := sortList(head)
	right = sortList(right)

	return merge(left, right)
}

func getMid(head *Node) *Node {
	slow := head
	fast := head.next

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func merge(l1, l2 *Node) *Node {
	dummy := &Node{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.data < l2.data {
			tail.next = l1
			l1 = l1.next
		} else {
			tail.next = l2
			l2 = l2.next
		}
		tail = tail.next
	}

	if l1 != nil {
		tail.next = l1
	}
	if l2 != nil {
		tail.next = l2
	}

	return dummy.next
}

func (l *SL) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("END")
}

func main() {
	list := &SL{}
	values := []int{4, 2, 1, 3}

	for _, v := range values {
		list.Insert(v)
	}

	fmt.Println("List before sorting : ")
	list.Display()

	list.head = sortList(list.head) // update head after sorting

	fmt.Println("List after sorting : ")
	list.Display()
}

// type Node struct {
// 	data int
// 	next *Node
// }

// type SL struct {
// 	head *Node
// }

// func (l *SL) Insert(value int) {
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

// func sortList(head *Node) *Node {
// 	if head == nil || head.next == nil {
// 		return head
// 	}

// 	mid := getMid(head)
// 	right := mid.next
// 	mid.next = nil

// 	left := sortList(head)
// 	right = sortList(right)

// 	return merge(left, right)
// }

// func getMid(head *Node) *Node {
// 	slow := head
// 	fast := head.next

// 	for fast != nil && fast.next != nil {
// 		slow = slow.next
// 		fast = fast.next.next
// 	}
// 	return slow
// }

// func merge(l1, l2 *Node) *Node {
// 	dummy := &Node{}
// 	tail := dummy

// 	for l1 != nil && l2 != nil {
// 		if l1.data < l2.data {
// 			tail.next = l1
// 			l1 = l1.next
// 		} else {
// 			tail.next = l2
// 			l2 = l2.next
// 		}
// 		tail = tail.next
// 	}

// 	if l1 != nil {
// 		tail.next = l1
// 	}
// 	if l2 != nil {
// 		tail.next = l2
// 	}

// 	return dummy.next
// }

// func (l *SL) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("END")
// }
