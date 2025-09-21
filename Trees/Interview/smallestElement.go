// Kth smallest element in BST
// Facebook Amazon and Google question

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

var count int

func smallest(root *Node, k int) int {
	count = 0
	node := helper(root, k)
	if node == nil {
		return -1
	}
	return node.data
}

func helper(root *Node, k int) *Node {
	if root == nil {
		return nil
	}

	left := helper(root.left, k)

	if left != nil {
		return left
	}

	count++

	if count == k {
		return root
	}

	return helper(root.right, k)
}

func main() {
	root := &Node{data: 3}
	root.left = &Node{data: 1}
	root.right = &Node{data: 4}
	root.left.right = &Node{data: 2}

	k := 3
	fmt.Println("smallest kth element is : ", smallest(root, k))
}
