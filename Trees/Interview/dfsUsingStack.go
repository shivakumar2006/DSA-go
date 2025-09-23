// DFS using stack

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func dfsStack(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{}
	stack = append(stack, root)

	for len(stack) > 0 {
		// pop from end
		n := len(stack) - 1
		current := stack[n]
		stack = stack[:n]

		fmt.Print(current.data, " ")

		if current.right != nil {
			stack = append(stack, current.right)
		}
		if current.left != nil {
			stack = append(stack, current.left)
		}
	}
}

func main() {
	root := &Node{data: 5}
	root.left = &Node{data: 4}
	root.right = &Node{data: 8}
	root.left.left = &Node{data: 11}
	root.right.left = &Node{data: 13}
	root.right.right = &Node{data: 4}

	fmt.Println("DFS traversal:")
	dfsStack(root)
}
