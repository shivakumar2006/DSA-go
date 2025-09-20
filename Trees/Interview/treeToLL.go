// Flatten Binary Tree to linked list
// Facebook question

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func flatten(root *Node) {
	if root == nil {
		return
	}

	current := root

	for current != nil {
		if current.left != nil {
			temp := current.left
			if temp.right != nil {
				temp = temp.right
			}

			temp.right = current.right
			current.right = current.left
			current.left = nil
		}
		current = current.right
	}
}

func DisplayFlatten(root *Node) {
	for root != nil {
		fmt.Print(root.data, " -> ")
		root = root.right
	}
	fmt.Println("END")
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 5}
	root.left.left = &Node{data: 3}
	root.left.right = &Node{data: 4}
	root.right.right = &Node{data: 6}

	flatten(root)
	DisplayFlatten(root)
}
