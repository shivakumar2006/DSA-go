// bubble sort in linkedlist

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

func (l *LinkedList) Length() int {
	count := 0
	temp := l.head
	for temp != nil {
		count++
		temp = temp.next
	}
	return count
}

func (l *LinkedList) bubbleSort() {
	n := l.Length()

	for row := 0; row < n-1; row++ {
		prev := l.head
		current := l.head
		for col := 0; col < n-row-1; col++ {
			next := current.next
			if current.data > next.data {
				if l.head == current {
					current.next = next.next
					next.next = current
					l.head = next
					prev = next
				} else {
					current.next = next.next
					next.next = current
					prev.next = next
					prev = next
				}
			} else {
				prev = current
				current = next
			}
		}
	}
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
	list := &LinkedList{}

	list.Insert(4)
	list.Insert(3)
	list.Insert(1)
	list.Insert(2)
	list.Insert(5)

	fmt.Print("Original list : ")
	list.Display()

	list.bubbleSort()

	fmt.Print("Sorted list : ")
	list.Display()
}

// func (l *LinkedList) bubbleSort() {
// 	n := l.Length()

// 	for row := 0; row < n-1; row++ {
// 		prev := l.head
// 		current := l.head
// 		for col := 0; col < n-row-1; col++ {
// 			next := current.next
// 			if current.data > next.data {
// 				if l.head == current {
// 					current.next = next.next
// 					next.next = current
// 					l.head = next
// 					prev = next
// 				} else {
// 					current.next = next.next
// 					next.next = current
// 					prev.next = next
// 					prev = next
// 				}
// 			} else {
// 				prev = current
// 				current = next
// 			}
// 		}
// 	}
// }

// func (l *LinkedList) bubbleSort() {
// 	n := l.Length()

// 	for row := 0; row < n-1; row++ {
// 		prev := l.head
// 		current := l.head
// 		for col := 0; col < n-row-1; col++ {
// 			next := current.next
// 			if current.data > next.data {
// 				if current == l.head {
// 					current.next = next.next
// 					next.next = current
// 					l.head = next
// 					prev = next
// 				} else {
// 					current.next = next.next
// 					next.next = current
// 					prev.next = next
// 					prev = next
// 				}
// 			} else {
// 				prev = current
// 				current = next
// 			}
// 		}
// 	}
// }

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

// func (l *LinkedList) Length() int {
// 	count := 0
// 	temp := l.head
// 	for temp != nil {
// 		count++
// 		temp = temp.next
// 	}

// 	return count
// }

// func (l *LinkedList) bubbleSort() {
// 	n := l.Length()

// 	for row := 0; row < n-1; row++ {
// 		prev := l.head
// 		current := l.head
// 		for col := 0; col < n-row-1; col++ {
// 			next := current.next
// 			if current.data > next.data {
// 				if current == l.head {
// 					current.next = next.next
// 					next.next = current
// 					l.head = next
// 					prev = next
// 				} else {
// 					current.next = next.next
// 					next.next = current
// 					prev.next = next
// 					prev = next
// 				}
// 			} else {
// 				// no swap just move forward
// 				prev = current
// 				current = next
// 			}
// 		}
// 	}
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

// 	list.Insert(4)
// 	list.Insert(5)
// 	list.Insert(2)
// 	list.Insert(1)
// 	list.Insert(3)

// 	fmt.Print("Original list : ")
// 	list.Display()

// 	list.bubbleSort()

// 	fmt.Print("sorted list : ")
// 	list.Display()
// }
