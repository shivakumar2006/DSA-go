// Binary tree path sum
// Facebook question \

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

var ans int // global maximum

func maxPathSum(root *Node) int {
	ans = math.MinInt // initialize to smallest unit
	helper(root)
	return ans
}

func helper(node *Node) int {
	if node == nil {
		return 0
	}

	// compute max path sum starting from left/right child
	left := helper(node.left)
	right := helper(node.right)

	// ignore negative paths
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}

	// path sum going through this node
	pathSum := left + right + node.data

	// update global maximum
	if pathSum > ans {
		ans = pathSum
	}

	// return the max single path
	if left > right {
		return left + node.data
	} else {
		return right + node.data
	}
}

func main() {
	root := &Node{data: -10}
	root.left = &Node{data: 9}
	root.right = &Node{data: 20}
	root.right.left = &Node{data: 15}
	root.right.right = &Node{data: 7}

	fmt.Println("Maximum Path Sum:", maxPathSum(root)) // Output: 42
}
