// maximum depth in binaryTree
// Google question

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

var depth int = 0

func depthTree(root *Node) int {
	MaxDepth(root)
	return depth
}

func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.left)
	right := MaxDepth(root.right)

	max := left + right

	if max > depth {
		depth = max
	}

	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}

	fmt.Println("the max depth is left side : ", depthTree(root))
}
