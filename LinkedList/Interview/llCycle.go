// Linked list cycle and find the length of the cycle as well
// Amazon and Microsoft question

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

func (l *LinkedList) hasCycle() bool {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return true
		}
	}
	return false
}

func (l *LinkedList) cycleLength() int {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			temp := slow.next
			length := 1
			for temp != slow {
				temp = temp.next
				length++
			}
			return length
		}
	}
	return 0
}

func (l *LinkedList) Display(limit int) {
	temp := l.head
	count := 0

	for temp != nil && count < limit {
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

	temp := list.head
	var thirdNode *Node
	for temp.next != nil {
		if temp.data == 3 {
			thirdNode = temp
		}
		temp = temp.next
	}
	temp.next = thirdNode

	if list.hasCycle() {
		fmt.Println("Cycle is present in the list...")
		length := list.cycleLength()
		fmt.Println("length of the cycle in the list : ", length)
	} else {
		fmt.Println("Cycle is not present in the list.")
	}

	list.Display(15)
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

// func (l *LinkedList) hasCycle() bool {
// 	fast := l.head
// 	slow := l.head

// 	for fast != nil && fast.next != nil {
// 		fast = fast.next.next
// 		slow = slow.next
// 		if fast == slow {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (l *LinkedList) cycleLength() int {
// 	fast := l.head
// 	slow := l.head

// 	for fast != nil && fast.next != nil {
// 		fast = fast.next.next
// 		slow = slow.next
// 		if fast == slow {
// 			// calculate the length
// 			length := 1
// 			temp := slow.next
// 			for temp != slow {
// 				temp = temp.next
// 				length++
// 			}
// 			return length
// 		}
// 	}
// 	return 0
// }

// func (l *LinkedList) Display() {
// 	temp := l.head
// 	for temp != nil {
// 		fmt.Print(temp.data, " -> ")
// 		temp = temp.next
// 	}
// 	fmt.Println("End")
// }

// func main() {
// 	list := &LinkedList{}

// 	list.Insert(1)
// 	list.Insert(2)
// 	list.Insert(3)
// 	list.Insert(4)
// 	list.Insert(5)
// 	list.Insert(6)

// 	// create a cycle for testing
// 	temp := list.head
// 	var thirdNode *Node
// 	for temp.next != nil {
// 		if temp.data == 3 {
// 			thirdNode = temp
// 		}
// 		temp = temp.next
// 	}
// 	temp.next = thirdNode // create the cycle

// 	if list.hasCycle() {
// 		fmt.Println("cycle is present in the list")
// 		length := list.cycleLength()
// 		fmt.Println("Cycle length : ", length)
// 	} else {
// 		fmt.Println("cycle is not present in the list")
// 	}
// }

// func (l *LinkedList) hasCycle() bool {
// 	fast := l.head
// 	slow := l.head

// 	for fast != nil && fast.next != nil {
// 		fast = fast.next.next
// 		slow = slow.next
// 		if fast == slow {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (l *LinkedList) cycleLength() int {
// 	fast := l.head
// 	slow := l.head

// 	for fast != nil && fast.next != nil {
// 		fast = fast.next.next
// 		slow = slow.next
// 		if fast == slow {
// 			// calculate the length
// 			temp := slow.next
// 			length := 1
// 			for temp != slow {
// 				temp = temp.next
// 				length++
// 			}
// 			return length
// 		}
// 	}
// 	return 0
// }
