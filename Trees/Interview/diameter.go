// Diameter in BinarySearch
// Google, Facebook and Amazon question

package main

import "math"

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

func height(root *Node) {
	if root == nil {
		return
	}

	leftHeight := height(root.left)
	rightHeight := height(root.right)

	dia := leftHeight - rightHeight + 1
	diameter = math.Max(diameter, dia)

	return math.Max(leftHeight, rightHeight) + 1
}
