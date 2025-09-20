// Invert Binary Tree

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func InvertTree(root *Node) *Node {
	if root == nil {
		return nil
	}

	Left := InvertTree(root.left)
	Right := InvertTree(root.right)

	root.left = Right
	root.right = Left

	return root
}

func Display(root *Node) {
	if root == nil {
		return
	}
	Display(root.left)

	fmt.Print(root.data, "")

	Display(root.right)
}

func DisplayLevelOrderTree(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			fmt.Println(current.data, " ")

			if current.left != nil {
				queue = append(queue, current.left)
			}
			if current.right != nil {
				queue = append(queue, current.right)
			}
		}
		fmt.Println()
	}
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}

	fmt.Println("Original tree: ")
	Display(root)
	fmt.Println()

	root = InvertTree(root)

	fmt.Println("Inverted tree : ")
	Display(root)
	fmt.Println()

	fmt.Print("Inverted tree : (level-order)")
	DisplayLevelOrderTree(root)
}
