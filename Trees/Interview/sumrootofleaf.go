// sum root of leaf numbers

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func sumRootToLeaf(root *Node) int {
	return helper(root, 0)
}

func helper(root *Node, current int) int {
	if root == nil {
		return 0
	}

	current = current*10 + root.data

	if root.left == nil && root.right == nil {
		return current
	}

	return helper(root.left, current) + helper(root.right, current)
}

func main() {
	root := &Node{data: 4}
	root.left = &Node{data: 9}
	root.left.left = &Node{data: 5}
	root.left.right = &Node{data: 1}
	root.right = &Node{data: 0}

	fmt.Println("Sum of root-to-leaf numbers:", sumRootToLeaf(root))
}

// func sumRootToLeaf(root *Node) int {
// 	return helper(root, 0)
// }

// func helper(root *Node, current int) int {
// 	if root == nil {
// 		return 0
// 	}

// 	current = current*10 + root.data

// 	if root.left == nil && root.right == nil {
// 		return current
// 	}

// 	return helper(root.left, current) + helper(root.right, current)
// }

// func sumRootToLeaf(root *Node) int {
// 	return helper(root, 0)
// }

// func helper(root *Node, current int) int {
// 	if root == nil {
// 		return 0
// 	}

// 	current = current*10 + root.data

// 	// if leaf node return the number
// 	if root.left == nil && root.right == nil {
// 		return current
// 	}

// 	return helper(root.left, current) + helper(root.right, current)
// }
