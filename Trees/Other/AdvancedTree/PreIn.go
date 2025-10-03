// Construct BinaryTree from preorder and inorder traversal

package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

var preIndex int

func buildTree(preorder, inorder []int) *Node {
	preIndex = 0
	maps := make(map[int]int)

	for i := 0; i < len(inorder); i++ {
		maps[inorder[i]] = i
	}

	return helper(preorder, inorder, 0, len(inorder)-1, maps)
}

func helper(preorder, inorder []int, left, right int, maps map[int]int) *Node {
	if left > right {
		return nil
	}

	rootVal := preorder[preIndex]
	preIndex++

	root := &Node{val: rootVal}

	if left == right {
		return root
	}

	inorderIndex := maps[rootVal]

	root.left = helper(preorder, inorder, left, inorderIndex-1, maps)
	root.right = helper(preorder, inorder, inorderIndex+1, right, maps)

	return root
}

func printInorder(root *Node) {
	if root == nil {
		return
	}

	printInorder(root.left)
	fmt.Print(root.val, " ")
	printInorder(root.right)
}

func printPreorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.val, " ")
	printPreorder(root.left)
	printPreorder(root.right)
}

func displayInorder(root *Node, level int) {
	if root == nil {
		return
	}

	displayInorder(root.right, level+1)
	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|\t")
		}
		fmt.Println("|----->", root.val)
	} else {
		fmt.Println(root.val)
	}
	displayInorder(root.left, level+1)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)

	fmt.Println("Inorder traversal of constructed tree : ")
	printInorder(root)
	fmt.Println()
	fmt.Println("inorder tree : ")
	displayInorder(root, 0)

	fmt.Println()
	fmt.Println("Preorder traversal of constructed tree : ")
	printPreorder(root)
	fmt.Println()

	fmt.Println("preorder tree : ")
	displayInorder(root, 0)
}

// type Node struct {
// 	val   int
// 	left  *Node
// 	right *Node
// }

// var preIndex int

// func buildTree(preorder, inorder []int) *Node {
// 	preIndex = 0
// 	maps := make(map[int]int)

// 	// store inorder position in map
// 	for i := 0; i < len(inorder); i++ {
// 		maps[inorder[i]] = i
// 	}

// 	return helper(preorder, inorder, 0, len(inorder)-1, maps)
// }

// func helper(preorder, inorder []int, left, right int, maps map[int]int) *Node {
// 	if left > right {
// 		return nil
// 	}

// 	rootVal := preorder[preIndex]
// 	preIndex++

// 	root := &Node{val: rootVal}

// 	// if no children
// 	if left == right {
// 		return root
// 	}

// 	// split inorder by root position
// 	inorderIndex := maps[rootVal]

// 	root.left = helper(preorder, inorder, left, inorderIndex-1, maps)
// 	root.right = helper(preorder, inorder, inorderIndex+1, right, maps)

// 	return root
// }
