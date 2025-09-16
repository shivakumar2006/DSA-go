// // Breadth First Search (BFS)

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func BFS(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Print(current.data, " ")

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func Display(node *Node, level int) {
	if node == nil {
		return
	}

	Display(node.right, level+1)

	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|\t")
		}
		fmt.Println("|---->", node.data)
	} else {
		fmt.Println(node.data)
	}

	Display(node.left, level+1)
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}

	fmt.Println("BFS traversal : ")
	BFS(root)
	fmt.Println()

	fmt.Println("\nDisplay tree")
	Display(root, 0)
}

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func BFS(root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	queue := []*Node{root}

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]

// 		fmt.Print(current.data, " ")

// 		if current.left != nil {
// 			queue = append(queue, current.left)
// 		}

// 		if current.right != nil {
// 			queue = append(queue, current.right)
// 		}
// 	}
// }

// func Display(node *Node, level int) {
// 	if node == nil {
// 		return
// 	}

// 	Display(node.left, level+1)

// 	if level != 0 {
// 		for i := 0; i < level-1; i++ {
// 			fmt.Print("|\t")
// 		}
// 		fmt.Println("|----->", node.data)
// 	} else {
// 		fmt.Println(node.data)
// 	}
// 	Display(node.right, level+1)
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.left.right = &Node{data: 5}
// 	root.right.left = &Node{data: 6}
// 	root.right.right = &Node{data: 7}

// 	fmt.Println("BFS traversal")
// 	BFS(root)
// 	fmt.Println()

// 	fmt.Println("\ndiaplay tree : ")
// 	Display(root, 0)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// // BFS traversal
// func BFS(root *Node) {
// 	if root == nil {
// 		return
// 	}

// 	// create queue slice to store nodes level by level
// 	queue := []*Node{root}

// 	for len(queue) > 0 {
// 		// Dequeue the first element
// 		current := queue[0]
// 		queue = queue[1:]

// 		// process the current node
// 		fmt.Print(current.data, " ")

// 		// Enqueue left child if exist
// 		if current.left != nil {
// 			queue = append(queue, current.left)
// 		}

// 		// Enqueue right child if exists
// 		if current.right != nil {
// 			queue = append(queue, current.right)
// 		}
// 	}
// }

// func prettyDisplay(node *Node, level int) {
// 	if node == nil {
// 		return
// 	}

// 	prettyDisplay(node.right, level+1)

// 	if level != 0 {
// 		for i := 0; i < level-1; i++ {
// 			fmt.Print("|\t")
// 		}
// 		fmt.Println("|----->", node.data)
// 	} else {
// 		fmt.Println(node.data)
// 	}
// 	prettyDisplay(node.left, level+1)
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.left.right = &Node{data: 5}
// 	root.right.left = &Node{data: 6}
// 	root.right.right = &Node{data: 7}

// 	fmt.Print("BFS Traversal : ")
// 	BFS(root)
// 	fmt.Println()

// 	fmt.Println("\nTree Structure:")
// 	prettyDisplay(root, 0)
// }
