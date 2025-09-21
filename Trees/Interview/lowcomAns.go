// Lowest common ancestor in binary tree
// Amazon question

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func lowestCommonAncestor(root *Node, p, q *Node) *Node {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.left, p, q)
	right := lowestCommonAncestor(root.right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return right
	}
	return left
}

func display(root *Node, level int) {
	if root == nil {
		return
	}

	display(root.right, level+1)

	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|\t")
		}
		fmt.Println("|------>", root.data)
	} else {
		fmt.Println(root.data)
	}
	display(root.left, level+1)
}

func main() {
	root := &Node{data: 3}
	root.left = &Node{data: 5}
	root.right = &Node{data: 1}
	root.left.left = &Node{data: 6}
	root.left.right = &Node{data: 2}
	root.right.left = &Node{data: 0}
	root.right.right = &Node{data: 8}
	root.left.right.left = &Node{data: 7}
	root.left.right.right = &Node{data: 4}

	p := root.left
	q := root.right

	ancestor := lowestCommonAncestor(root, p, q)
	fmt.Println("Lowest common ancestor is : ", ancestor.data)
	fmt.Println()

	fmt.Println("Tree : ")
	display(root, 0)
}
