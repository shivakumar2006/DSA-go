// // Validate Binary search tree

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func isValidate(root *Node) bool {
	return helper(root, nil, nil)
}

func helper(root *Node, low *int, high *int) bool {
	if root == nil {
		return true
	}

	if low != nil && root.data <= *low {
		return false
	}

	if high != nil && root.data >= *high {
		return false
	}

	leftTree := helper(root.left, low, &root.data)
	rightTree := helper(root.right, &root.data, high)

	return leftTree && rightTree
}

func main() {
	root := &Node{data: 2}
	root.left = &Node{data: 1}
	root.right = &Node{data: 3}

	fmt.Println(isValidate(root))
}

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func isValidBST(root *Node) bool {
// 	return helper(root, nil, nil)
// }

// func helper(root *Node, low *int, high *int) bool {
// 	if root == nil {
// 		return true
// 	}

// 	if low != nil && root.data <= *low {
// 		return false
// 	}

// 	if high != nil && root.data >= *high {
// 		return false
// 	}

// 	leftTree := helper(root.left, low, &root.data)
// 	rightTree := helper(root.right, &root.data, high)

// 	return leftTree && rightTree
// }

// func main() {
// 	root := &Node{data: 2}
// 	root.left = &Node{data: 1}
// 	root.right = &Node{data: 3}

// 	fmt.Println(isValidBST(root))
// }
