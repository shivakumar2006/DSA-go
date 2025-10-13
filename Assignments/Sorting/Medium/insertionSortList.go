// Insertion Sort List
// https://leetcode.com/problems/insertion-sort-list/description/

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type ISL struct {
	head *Node
}

func (ll *ISL) Insert(value int) {
	newNode := &Node{data: value}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	temp := ll.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
}

func insertionSortList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	dummy := &Node{}
	current := head

	for current != nil {
		prev := dummy
		next := current.next

		// find the right place to insert current node
		for prev.next != nil && prev.next.data < current.data {
			prev = prev.next
		}

		// insert current between prev and prev.next
		current.next = prev.next
		prev.next = current

		current = next
	}
	return dummy.next
}

func (ll *ISL) Display() {
	temp := ll.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("END")
}

func main() {
	ll := &ISL{}
	values := []int{4, 2, 1, 3}

	for _, v := range values {
		ll.Insert(v)
	}

	fmt.Println("Before sorting : ")
	ll.Display()

	ll.head = insertionSortList(ll.head)

	fmt.Println("After sorting : ")
	ll.Display()
}
