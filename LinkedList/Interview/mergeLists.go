// // // merge two sorted linked list

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

func mergeLists(l1, l2 *LinkedList) *LinkedList {
	first := l1.head
	second := l2.head

	ans := &LinkedList{}

	for first != nil && second != nil {
		if first.data < second.data {
			ans.Insert(first.data)
			first = first.next
		} else {
			ans.Insert(second.data)
			second = second.next
		}
	}

	// if there is any element is present in any list after completing ont list then add rest of the nodes or elemets from the existing list
	for first != nil {
		ans.Insert(first.data)
		first = first.next
	}

	for second != nil {
		ans.Insert(second.data)
		second = second.next
	}

	return ans
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.next
	}
	fmt.Println("End")
}

func main() {
	list1 := &LinkedList{}
	list2 := &LinkedList{}

	list1.Insert(1)
	list1.Insert(3)
	list1.Insert(5)

	list2.Insert(2)
	list2.Insert(4)
	list2.Insert(6)
	list2.Insert(7)

	list1.Display()
	list2.Display()

	ans := mergeLists(list1, list2)
	ans.Display()
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
// 	temp = newNode
// }

// func mergeLists(l1, l2 *LinkedList) *LinkedList {
// 	first := l1.head
// 	second := l2.head

// 	ans := &LinkedList{}

// 	for first != nil && second != nil {
// 		if first.data < second.data {
// 			ans.Insert(first.data)
// 			first = first.next
// 		} else {
// 			ans.Insert(second.data)
// 			second = second.next
// 		}
// 	}

// 	// add rest of the nodes in the list
// 	for first != nil {
// 		ans.Insert(first.data)
// 		first = first.next
// 	}

// 	for second != nil {
// 		ans.Insert(second.data)
// 		second = second.next
// 	}

// 	return ans
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
// 	list1 := &LinkedList{}
// 	list2 := &LinkedList{}

// 	list1.Insert(1)
// 	list1.Insert(3)
// 	list1.Insert(5)

// 	list2.Insert(2)
// 	list2.Insert(4)
// 	list2.Insert(6)
// 	list2.Insert(7)

// 	list1.Display()
// 	list2.Display()

// 	ans := mergeLists(list1, list2)
// 	ans.Display()
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
// // 	temp = newNode
// // }

// // func mergeLists(l1, l2 *LinkedList) *LinkedList {
// // 	first := l1.head
// // 	second := l2.head

// // 	ans := &LinkedList{}

// // 	for first != nil && second != nil {
// // 		if first.data < second.data {
// // 			ans.Insert(first.data)
// // 			first = first.next
// // 		} else {
// // 			ans.Insert(second.data)
// // 			second = second.next
// // 		}
// // 	}

// // 	for first != nil {
// // 		ans.Insert(first.data)
// // 		first = first.next
// // 	}

// // 	for second != nil {
// // 		ans.Insert(second.data)
// // 		second = second.next
// // 	}

// // 	return ans
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
// // 	list1 := &LinkedList{}
// // 	list2 := &LinkedList{}

// // 	list1.Insert(1)
// // 	list1.Insert(3)
// // 	list1.Insert(5)

// // 	list2.Insert(2)
// // 	list2.Insert(4)
// // 	list2.Insert(6)

// // 	list1.Display()
// // 	list2.Display()

// // 	ans := mergeLists(list1, list2)
// // 	ans.Display()
// // }
