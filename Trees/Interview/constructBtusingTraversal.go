// // Construct BinaryTree from Pre-order and in-order traversal

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func construct(preorder []int, inorder []int) *Node {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootValue := preorder[0]
	root := &Node{data: rootValue}

	var rootIndex int
	for i, v := range inorder {
		if v == rootValue {
			rootIndex = i
			break
		}
	}

	root.left = construct(preorder[1:1+rootIndex], inorder[:rootIndex+1])
	root.right = construct(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root
}

func displayPreorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.data, " ")
	displayPreorder(root.left)
	displayPreorder(root.right)
}

func displayInorder(root *Node) {
	if root == nil {
		return
	}

	displayInorder(root.left)
	fmt.Print(root.data, " ")
	displayInorder(root.right)
}

func Display(root *Node, level int) {
	if root == nil {
		return
	}

	Display(root.right, level+1)

	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|\t")
		}
		fmt.Println("|----->", root.data)
	} else {
		fmt.Println(root.data)
	}
	Display(root.left, level+1)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := construct(preorder, inorder)

	fmt.Println("preorder traversal of tree : ")
	displayPreorder(root)
	fmt.Println()

	fmt.Println("inorder traversal of tree : ")
	displayInorder(root)
	fmt.Println()

	fmt.Println("Display the tree : ")
	Display(root, 0)
}

// func construct(preorder []int, inorder []int) *Node {
// 	if len(preorder) == 0 || len(inorder) == 0 {
// 		return nil
// 	}

// 	rootValue := preorder[0]
// 	root := &Node{data: rootValue}

// 	var rootIndex int
// 	for i, v := range inorder {
// 		if v == rootValue {
// 			rootIndex = i
// 			break
// 		}
// 	}

// 	root.left = construct(preorder[1:1+rootIndex], inorder[:rootIndex+1])
// 	root.right = construct(preorder[1+rootIndex:], inorder[rootIndex+1:])

// 	return root
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func buildTree(preorder []int, inorder []int) *Node {
// 	if len(preorder) == 0 || len(inorder) == 0 {
// 		return nil
// 	}

// 	// the first element of preorder is a root
// 	rootVal := preorder[0]
// 	root := &Node{data: rootVal}

// 	// find the index of the root in preorder
// 	var rootIndex int
// 	for i, v := range inorder {
// 		if v == rootVal {
// 			rootIndex = i
// 			break
// 		}
// 	}

// 	// construct left and right subtree recursively
// 	root.left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
// 	root.right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])

// 	return root
// }

// func displayInorder(root *Node) {
// 	if root == nil {
// 		return
// 	}
// 	displayInorder(root.left)
// 	fmt.Print(root.data, " ")
// 	displayInorder(root.right)
// }

// func displatPreorder(root *Node) {
// 	if root == nil {
// 		return
// 	}
// 	fmt.Print(root.data, " ")
// 	displayInorder(root.left)
// 	displayInorder(root.right)
// }

// // Pretty display of the tree
// func displayTree(root *Node, level int) {
// 	if root == nil {
// 		return
// 	}

// 	displayTree(root.right, level+1)

// 	if level != 0 {
// 		for i := 0; i < level-1; i++ {
// 			fmt.Print("|\t")
// 		}
// 		fmt.Println("|---->", root.data)
// 	} else {
// 		fmt.Println(root.data)
// 	}

// 	displayTree(root.left, level+1)
// }

// func main() {
// 	preorder := []int{3, 9, 20, 15, 7}
// 	inorder := []int{9, 3, 15, 20, 7}

// 	root := buildTree(preorder, inorder)

// 	fmt.Print("Inorder of constructed tree: ")
// 	displayInorder(root)
// 	fmt.Println()

// 	fmt.Print("Preorder of constructed tree: ")
// 	displatPreorder(root)
// 	fmt.Println()

// 	fmt.Println("Pretty display of the constructed tree:")
// 	displayTree(root, 0)
// }
