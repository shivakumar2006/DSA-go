// Binary tree path sum
// Facebook question \

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func pathSum(root *Node, target int) [][]int {
	var result [][]int
	var currentPath []int
	helper(root, target, currentPath, &result)
	return result
}

func helper(root *Node, target int, path []int, result *[][]int) {
	if root == nil {
		return
	}

	// add current node to path
	path = append(path, root.data)

	// check if it's a. leaf and the sum equals target.
	if root.left == nil && root.right == nil && root.data == target {
		// copy the path and append to result
		temp := make([]int, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// recursed left and right with reduced target
	helper(root.left, target-root.data, path, result)
	helper(root.right, target-root.data, path, result)
}

func main() {
	root := &Node{data: -10}
	root.left = &Node{data: 9}
	root.right = &Node{data: 20}
	root.right.left = &Node{data: 15}
	root.right.right = &Node{data: 7}

	target := 25
	path := pathSum(root, target)
	fmt.Println("path with sum", target, ":")
	for _, p := range path {
		fmt.Println(p)
	}
}
