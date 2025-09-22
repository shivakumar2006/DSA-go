// // Path sum
// // Amazon question

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func pathSum(root *Node, sum int) bool {
	if root == nil {
		return false
	}

	if root.data == sum && root.left == nil && root.right == nil {
		return true
	}

	return pathSum(root.left, sum-root.data) || pathSum(root.right, sum-root.data)
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
		fmt.Println("|----->", root.data)
	} else {
		fmt.Println(root.data)
	}
	display(root.left, level+1)
}

func main() {
	root := &Node{data: 5}
	root.left = &Node{data: 4}
	root.right = &Node{data: 8}
	root.left.left = &Node{data: 11}
	root.left.left.left = &Node{data: 7}
	root.left.left.right = &Node{data: 2}
	root.right.left = &Node{data: 13}
	root.right.right = &Node{data: 4}
	root.right.right.right = &Node{data: 1}

	fmt.Println(pathSum(root, 22))
	fmt.Println(pathSum(root, 26))
	fmt.Println(pathSum(root, 18))

	fmt.Println("Display tree : ")
	display(root, 0)
}

// func pathSum(root *Node, sum int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if root.data == sum && root.left == nil && root.right == nil {
// 		return true
// 	}

// 	return pathSum(root.left, sum-root.data) || pathSum(root.right, sum-root.data)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func pathSum(root *Node, sum int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if root.data == sum && root.left == nil && root.right == nil {
// 		return true
// 	}

// 	return pathSum(root.left, sum-root.data) || pathSum(root.right, sum-root.data)
// }

// func display(root *Node, level int) {
// 	if root == nil {
// 		return
// 	}

// 	display(root.right, level+1)
// 	if level != 0 {
// 		for i := 0; i < level-1; i++ {
// 			fmt.Print("|\t")
// 		}
// 		fmt.Println("|----->", root.data)
// 	} else {
// 		fmt.Println(root.data)
// 	}
// 	display(root.left, level+1)
// }

// func main() {
// 	root := &Node{data: 5}
// 	root.left = &Node{data: 4}
// 	root.right = &Node{data: 8}

// 	// left subtree
// 	root.left.left = &Node{data: 11}
// 	root.left.left.left = &Node{data: 7}
// 	root.left.left.right = &Node{data: 2}

// 	// right subtree
// 	root.right.left = &Node{data: 13}
// 	root.right.right = &Node{data: 4} // <-- you wrote 4 here
// 	root.right.right.right = &Node{data: 1}

// 	fmt.Println(pathSum(root, 22)) // true (5->4->11->2)
// 	fmt.Println(pathSum(root, 26)) // true (5->8->13)
// 	fmt.Println(pathSum(root, 18)) // false

// 	fmt.Println("Display tree : ")
// 	display(root, 0)
// }
