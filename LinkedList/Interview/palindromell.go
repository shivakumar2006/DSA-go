// // // Palindrome Linked list
// // // Google, Microsoft, Facebook, Linkedin, Amazon and Apple question

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

func isPalindrome(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	// find mid
	mid := middleNode(head)

	// check second part of the list
	secondHead := reverseList(mid)

	p1 := head
	p2 := secondHead
	for p2 != nil {
		if p1.data != p2.data {
			reverseList(secondHead)
			return false
		}
		p1 = p1.next
		p2 = p2.next
	}
	reverseList(secondHead)
	return true
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
	list.Insert(3)
	list.Insert(2)
	list.Insert(1)

	list.Display()

	if isPalindrome(list.head) {
		fmt.Println("current list is palindrome...")
	} else {
		fmt.Println("list is not a palindrome")
	}

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

// func isPalindrome(head *Node) bool {
// 	if head != nil || head.next != nil {
// 		return true
// 	}

// 	// find middle
// 	mid := middleNode(head)

// 	// second half
// 	secondHead := reverseList(mid)

// 	p1 := head
// 	p2 := secondHead
// 	for p2 != nil {
// 		if p1.data != p2.data {
// 			reverseList(secondHead)
// 			return false
// 		}
// 		p1 = p1.next
// 		p2 = p2.next
// 	}

// 	reverseList(secondHead)
// 	return true
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
// 	list.Insert(3)
// 	list.Insert(2)
// 	list.Insert(1)

// 	fmt.Print("Original list : ")
// 	list.Display()

// 	if isPalindrome(list.head) {
// 		fmt.Println("current list is palindrome")
// 	} else {
// 		fmt.Println("not a palindrome")
// 	}

// 	fmt.Print("List after check : ")
// 	list.Display()
// }

// // package main

// // import "fmt"

// // type Node struct {
// // 	data int
// // 	next *Node
// // }

// // type LinkedList struct {
// // 	head *Node
// // }

// // func (l *LinkedList) Insert(value int) {
// // 	newNode := &Node{data: value}

// // 	if l.head == nil {
// // 		l.head = newNode
// // 		return
// // 	}

// // 	temp := l.head
// // 	for temp.next != nil {
// // 		temp = temp.next
// // 	}

// // 	temp.next = newNode
// // }

// // func reverseList(head *Node) *Node {
// // 	var prev *Node
// // 	current := head

// // 	for current != nil {
// // 		next := current.next
// // 		current.next = prev
// // 		prev = current
// // 		current = next
// // 	}
// // 	return prev
// // }

// // func middleNode(head *Node) *Node {
// // 	slow := head
// // 	fast := head

// // 	for fast != nil && fast.next != nil {
// // 		slow = slow.next
// // 		fast = fast.next.next
// // 	}

// // 	return slow
// // }

// // func isPalindrome(head *Node) bool {
// // 	if head == nil || head.next == nil {
// // 		return true
// // 	}

// // 	// find middle
// // 	mid := middleNode(head)

// // 	// reverse second half
// // 	secondHead := reverseList(mid)

// // 	// compare
// // 	p1 := head
// // 	p2 := secondHead
// // 	for p2 != nil {
// // 		if p1.data != p2.data {
// // 			// restore before returning
// // 			reverseList(secondHead)
// // 			return false
// // 		}
// // 		p1 = p1.next
// // 		p2 = p2.next
// // 	}

// // 	// restore list
// // 	reverseList(secondHead)
// // 	return true
// // }

// // func (l *LinkedList) Display() {
// // 	temp := l.head
// // 	for temp != nil {
// // 		fmt.Print(temp.data, " -> ")
// // 		temp = temp.next
// // 	}
// // 	fmt.Println("End")
// // }

// // func main() {
// // 	list := &LinkedList{}

// // 	list.Insert(1)
// // 	list.Insert(2)
// // 	list.Insert(3)
// // 	list.Insert(3)
// // 	list.Insert(2)
// // 	list.Insert(1)

// // 	fmt.Print("Original list : ")
// // 	list.Display()

// // 	if isPalindrome(list.head) {
// // 		fmt.Println("Palindrome linkedlist...")
// // 	} else {
// // 		fmt.Println("not palindrome")
// // 	}

// // 	fmt.Print("List after check (restores) : ")
// // 	list.Display()
// // }
