// Diameter in BinarySearch
// Google, Facebook and Amazon question

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

var diameter int = 0

func diameterOfBinaryTree(root *Node) int {
	height(root)
	return diameter
}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.left)
	rightHeight := height(root.right)

	dia := leftHeight + rightHeight

	if dia > diameter {
		diameter = dia
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}

	fmt.Println("Diameter:", diameterOfBinaryTree(root)) // should print 3
}

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// var diameter int = 0

// func diameterOfBinaryTree(root *Node) int {
// 	height(root)
// 	return diameter
// }

// func height(root *Node) int {
// 	if root == nil {
// 		return 0
// 	}

// 	leftHeight := height(root.left)
// 	rightHeight := height(root.right)

// 	dia := leftHeight + rightHeight

// 	if dia > diameter {
// 		diameter = dia
// 	}

// 	if leftHeight > rightHeight {
// 		return leftHeight + 1
// 	}
// 	return rightHeight + 1
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// var diameter int = 0

// func diameterOfBinaryTree(root *Node) int {
// 	height(root)
// 	return diameter
// }

// func height(root *Node) int {
// 	if root == nil {
// 		return 0
// 	}

// 	leftHeight := height(root.left)
// 	rightHeight := height(root.right)

// 	dia := leftHeight + rightHeight

// 	if dia > diameter {
// 		diameter = dia
// 	}

// 	// return the height as in your last line but with int max
// 	if leftHeight > rightHeight {
// 		return leftHeight + 1
// 	}
// 	return rightHeight + 1
// }
