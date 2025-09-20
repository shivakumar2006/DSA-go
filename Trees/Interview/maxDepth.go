// maximum depth in binaryTree
// Google question

package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	Left := maxDepth(root.left)
	Right := maxDepth(root.right)

	var depth float64 = math.Max(float64(Left), float64(Right)) + 1

	return int(depth)
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}

	fmt.Println("the max depth is : ", maxDepth(root))
}
