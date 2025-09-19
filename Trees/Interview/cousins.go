// // Find cousins in Binary tree

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func isCousins(node *Node, x, y int) bool {
	xx := findNode(node, x)
	yy := findNode(node, y)

	if xx == nil && yy == nil {
		return false
	}

	return (level(node, xx, 0) == level(node, yy, 0)) && !isSiblings(node, xx, yy)
}

func findNode(node *Node, x int) *Node {
	if node == nil {
		return nil
	}

	if node.data == x {
		return node
	}

	n := findNode(node.left, x)
	if n != nil {
		return n
	}
	return findNode(node.right, x)
}

func isSiblings(node *Node, x, y *Node) bool {
	if node == nil {
		return false
	}

	if (node.left == x && node.right == y) || (node.left == y && node.right == x) {
		return true
	}
	return isSiblings(node.left, x, y) || isSiblings(node.right, x, y)
}

func level(node *Node, x *Node, lev int) int {
	if node == nil {
		return -1
	}

	if node == x {
		return lev
	}

	l := level(node.left, x, lev+1)
	if l != -1 {
		return l
	}
	return level(node.right, x, lev+1)
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.right.right = &Node{data: 5}

	fmt.Println(isCousins(root, 4, 5)) // true (4 and 5 are cousins)
	fmt.Println(isCousins(root, 4, 3)) // false
}

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func isCousins(root *Node, x, y int) bool {
// 	xx := findNode(root, x)
// 	yy := findNode(root, y)

// 	if xx == nil && yy == nil {
// 		return false
// 	}

// 	return (level(root, xx, 0) == level(root, yy, 0) && !isSiblings(root, xx, yy))
// }

// func findNode(root *Node, x int) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	if root.data == x {
// 		return root
// 	}
// 	if n := findNode(root.left, x); n != nil {
// 		return n
// 	}
// 	return findNode(root.right, x)
// }

// func isSiblings(node *Node, x, y *Node) bool {
// 	if node == nil {
// 		return false
// 	}

// 	if (node.left == x && node.right == y) || (node.left == y && node.right == x) {
// 		return true
// 	}
// 	return isSiblings(node.left, x, y) || isSiblings(node.right, x, y)
// }

// func level(root *Node, x *Node, lev int) int {
// 	if root == nil {
// 		return -1
// 	}

// 	if root == x {
// 		return lev
// 	}

// 	l := level(root.left, x, lev+1)
// 	if l != -1 {
// 		return l
// 	}
// 	return level(root.right, x, lev+1)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// // main function check if x and y are cousins
// func isCousins(root *Node, x, y int) bool {
// 	xx := findNode(root, x)
// 	yy := findNode(root, y)

// 	if xx == nil && yy == nil {
// 		return false
// 	}

// 	return (level(root, xx, 0) == level(root, yy, 0) && !isSiblings(root, xx, yy))
// }

// func findNode(node *Node, x int) *Node {
// 	if node == nil {
// 		return nil
// 	}

// 	if node.data == x {
// 		return node
// 	}
// 	if n := findNode(node.left, x); n != nil {
// 		return n
// 	}
// 	return findNode(node.right, x)
// }

// func isSiblings(node *Node, x, y *Node) bool {
// 	if node == nil {
// 		return false
// 	}

// 	if (node.left == x && node.right == y) || (node.left == y && node.right == x) {
// 		return true
// 	}
// 	return isSiblings(node.left, x, y) || isSiblings(node.right, x, y)
// }

// func level(node *Node, x *Node, lev int) int {
// 	if node == nil {
// 		return 0 // not found
// 	}

// 	if node == x {
// 		return lev
// 	}

// 	l := level(node.left, x, lev+1)
// 	if l != 0 {
// 		return l
// 	}
// 	return level(node.right, x, lev+1)
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.right.right = &Node{data: 5}

// 	fmt.Println(isCousins(root, 4, 5)) // true (4 and 5 are cousins)
// 	fmt.Println(isCousins(root, 4, 3)) // false
// }
