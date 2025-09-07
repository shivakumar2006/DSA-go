// // Reverse K Nodes

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

func reverseKGroup(head *Node, k int) *Node {
	if k <= 1 || head == nil {
		return head
	}

	var prev *Node = nil
	current := head

	for {
		last := prev
		newEnd := current

		// reverse between left and right
		var next *Node = current.next
		for i := 0; current != nil && i < k; i++ {
			current.next = prev
			prev = current
			current = next
			if next != nil {
				next = next.next
			}
		}

		if last != nil {
			last.next = prev
		} else {
			head = prev
		}

		newEnd.next = current

		if current == nil {
			break
		}

		prev = newEnd
	}
	return head
}

func reverseAlternateKGroup(head *Node, k int) *Node {
	if k <= 1 || head == nil {
		return head
	}

	var prev *Node = nil
	current := head

	for current != nil {
		last := prev
		newEnd := current

		// reverse between left and right
		var next *Node = current.next
		for i := 0; current != nil && i < k; i++ {
			current.next = prev
			prev = current
			current = next
			if next != nil {
				next = next.next
			}
		}

		if last != nil {
			last.next = prev
		} else {
			head = prev
		}

		newEnd.next = current

		// skip next k node
		for i := 0; current != nil && i < k; i++ {
			prev = current
			current = current.next
		}

	}
	return head
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
	list.Insert(6)
	list.Insert(7)

	fmt.Print("Original list : ")
	list.Display()

	list.head = reverseKGroup(list.head, 3)

	fmt.Print("reverse k node list : ")
	list.Display()

	list.head = reverseAlternateKGroup(list.head, 3)

	fmt.Println("reverse alterante k node list : ")
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

// func reverseKGroup(head *Node, k int) *Node {
// 	if k <= 1 || head == nil {
// 		return head
// 	}

// 	var prev *Node = nil
// 	current := head

// 	for {
// 		last := prev
// 		newEnd := current

// 		// reverse between left and right
// 		var next *Node = current.next
// 		for i := 0; current != nil && i < k; i++ {
// 			current.next = prev
// 			prev = current
// 			current = next
// 			if next != nil {
// 				next = next.next
// 			}
// 		}

// 		if last != nil {
// 			last.next = prev
// 		} else {
// 			head = prev
// 		}

// 		newEnd.next = current

// 		if current == nil {
// 			break
// 		}

// 		prev = newEnd
// 	}
// 	return head
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
// 	list.Insert(6)
// 	list.Insert(7)

// 	fmt.Print("Original list : ")
// 	list.Display()

// 	list.head = reverseKGroup(list.head, 3)

// 	fmt.Print("Reversed in K Groups: ")
// 	list.Display()
// }
