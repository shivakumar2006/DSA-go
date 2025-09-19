// // Populating next right pointers in each node
// // Amazon London 2022 question

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}

	Connect(root)
	Display(root)
}

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// 	next  *Node
// }

// func Connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	leftMost := root

// 	for leftMost.left != nil {
// 		current := leftMost
// 		for current != nil {
// 			current.left.next = current.right
// 			if current.next != nil {
// 				current.right.next = current.next.left
// 			}
// 			current = current.next
// 		}
// 		leftMost = leftMost.left
// 	}
// 	return root
// }

// func Display(root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	result := []interface{}{}
// 	level := root

// 	for level != nil {
// 		current := level
// 		for current != nil {
// 			result = append(result, current.data)
// 			current = current.next
// 		}
// 		result = append(result, "#")
// 		level = level.left
// 	}
// 	fmt.Println(result)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// 	next  *Node
// }

// func Connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	leftMost := root

// 	for leftMost.left != nil {
// 		current := leftMost
// 		for current != nil {
// 			current.left.next = current.right
// 			if current.next != nil {
// 				current.right.next = current.next.left
// 			}
// 			current = current.next
// 		}
// 		leftMost = leftMost.left
// 	}
// 	return root
// }

// func Display(root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	result := []interface{}{}
// 	level := root

// 	for level != nil {
// 		current := level
// 		for current != nil {
// 			result = append(result, current.data)
// 			current = current.next
// 		}
// 		result = append(result, "#")
// 		level = level.left
// 	}
// 	fmt.Println(result)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// 	next  *Node
// }

// func Connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	leftMost := root

// 	for leftMost.left != nil {
// 		current := leftMost
// 		for current != nil {
// 			current.left.next = current.right
// 			if current.next != nil {
// 				current.right.next = current.next.left
// 			}
// 			current = current.next
// 		}
// 		leftMost = leftMost.left
// 	}
// 	return root
// }

// func Display(root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	result := []interface{}{}
// 	level := root

// 	for level != nil {
// 		current := level
// 		for current != nil {
// 			result = append(result, current.data)
// 			current = current.next
// 		}
// 		result = append(result, "#")
// 		level = level.left
// 	}
// 	fmt.Println(result)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// 	next  *Node
// }

// func Connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	leftMost := root

// 	for leftMost.left != nil {
// 		var current *Node = leftMost
// 		for current != nil {
// 			current.left.next = current.right
// 			if current.next != nil {
// 				current.right.next = current.next.left
// 			}
// 			current = current.next
// 		}
// 		leftMost = leftMost.left
// 	}
// 	return root
// }

// func Display(root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	result := []interface{}{}
// 	level := root

// 	for level != nil {
// 		current := level
// 		for current != nil {
// 			result = append(result, current.data)
// 			current = current.next
// 		}
// 		result = append(result, "#")
// 		level = level.left
// 	}
// 	fmt.Println(result)
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.left.right = &Node{data: 5}
// 	root.right.left = &Node{data: 6}
// 	root.right.right = &Node{data: 7}

// 	Connect(root)
// 	Display(root)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// 	next  *Node
// }

// func Connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	leftMost := root

// 	for leftMost.left != nil {
// 		var current *Node = leftMost
// 		for current != nil {
// 			current.left.next = current.right
// 			if current.next != nil {
// 				current.right.next = current.next.left
// 			}
// 			current = current.next
// 		}
// 		leftMost = leftMost.left
// 	}
// 	return root
// }

// func Display(root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	result := []interface{}{}
// 	level := root
// 	for level != nil {
// 		current := level
// 		for current != nil {
// 			result = append(result, current.data)
// 			current = current.next
// 		}
// 		result = append(result, "#") // end of this level
// 		level = level.left           // move to the next level
// 	}
// 	fmt.Println(result)
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.left.right = &Node{data: 5}
// 	root.right.left = &Node{data: 6}
// 	root.right.right = &Node{data: 7}

// 	Connect(root)
// 	Display(root)
// }
