// Binary Search Tree

package main

import "fmt"

type Node struct {
	data   int
	left   *Node
	right  *Node
	height int
}

type BinarySearchTree struct {
	root *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// create a newBST
func NewBST() *BinarySearchTree {
	return &BinarySearchTree{root: nil}
}

func (bst *BinarySearchTree) height(node *Node) int {
	if node == nil {
		return -1
	}

	return node.height
}

// Insert public
func (bst *BinarySearchTree) Insert(value int) {
	bst.root = bst.insert(bst.root, value)
}

// Insert private
func (bst *BinarySearchTree) insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{data: value}
	}
	if value < node.data {
		node.left = bst.insert(node.left, value)
	} else if value > node.data {
		node.right = bst.insert(node.right, value)
	}
	node.height = max(bst.height(node.left), bst.height(node.right)) + 1
	return node
}

// check whether the tree is balanced or not
func (bst *BinarySearchTree) Balanced() bool {
	return bst.balanced(bst.root)
}

func (bst *BinarySearchTree) balanced(node *Node) bool {
	if node == nil {
		return true
	}

	leftHeight := bst.height(node.left)
	rightHeight := bst.height(node.right)
	return abs(leftHeight-rightHeight) <= 1 && bst.balanced(node.left) && bst.balanced(node.right)
}

func (bst *BinarySearchTree) Display() {
	bst.display(bst.root, "Root Node: ")
}

func (bst *BinarySearchTree) display(node *Node, details string) {
	if node == nil {
		return
	}
	fmt.Println(details, node.data)
	bst.display(node.left, fmt.Sprintf("Left child of %d: ", node.data))
	bst.display(node.right, fmt.Sprintf("Right child of %d: ", node.data))
}

func main() {
	tree := NewBST()
	value := []int{10, 5, 15, 3, 7, 12, 20}
	for _, v := range value {
		tree.Insert(v)
	}

	tree.Display()

	fmt.Println("Balanced: ", tree.Balanced())
}
