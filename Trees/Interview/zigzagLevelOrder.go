// Binary tree Zigzag level order traversal
// Google, Microsoft and Amazon question...

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func ZigzagLevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	var result [][]int
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if leftToRight {
				level[i] = current.data
			} else {
				level[levelSize-1-i] = current.data
			}

			if current.left != nil {
				queue = append(queue, current.left)
			}
			if current.right != nil {
				queue = append(queue, current.right)
			}
		}
		result = append(result, level)
		leftToRight = !leftToRight
	}
	return result
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}

	fmt.Println("Zigzag level order traversal : ")
	result := ZigzagLevelOrder(root)
	fmt.Println(result)
}

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func ZigzagLevelOrder(root *Node) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	queue := []*Node{root}
// 	var result [][]int
// 	leftToRight := true

// 	for len(queue) > 0 {
// 		levelSize := len(queue)
// 		level := make([]int, levelSize)

// 		for i := 0; i < levelSize; i++ {
// 			current := queue[0]
// 			queue = queue[1:]

// 			if leftToRight {
// 				level[i] = current.data
// 			} else {
// 				level[levelSize-1-i] = current.data
// 			}

// 			if current.left != nil {
// 				queue = append(queue, current.left)
// 			}
// 			if current.right != nil {
// 				queue = append(queue, current.right)
// 			}
// 		}
// 		result = append(result, level)
// 		leftToRight = !leftToRight
// 	}
// 	return result
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func ZigzagLevelOrder(root *Node) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	queue := []*Node{root}
// 	var result [][]int
// 	leftToRight := true // toggle direction

// 	for len(queue) > 0 {
// 		levelSize := len(queue)
// 		level := make([]int, levelSize)

// 		for i := 0; i < levelSize; i++ {
// 			current := queue[0]
// 			queue = queue[1:]

// 			if leftToRight {
// 				level[i] = current.data
// 			} else {
// 				level[levelSize-1-i] = current.data
// 			}

// 			if current.left != nil {
// 				queue = append(queue, current.left)
// 			}
// 			if current.right != nil {
// 				queue = append(queue, current.right)
// 			}
// 		}
// 		result = append(result, level)
// 		leftToRight = !leftToRight
// 	}
// 	return result
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func ZigzagLevelOrder(root *Node) [][]int {
// 	if root == nil {
// 		return nil
// 	}

// 	queue := []*Node{root}
// 	var result [][]int
// 	leftToRight := true // toggle direction

// 	for len(queue) > 0 {
// 		levelSize := len(queue)
// 		level := make([]int, levelSize)

// 		for i := 0; i < levelSize; i++ {
// 			current := queue[0]
// 			queue = queue[1:]

// 			// position depends on direction
// 			if leftToRight {
// 				level[i] = current.data
// 			} else {
// 				level[levelSize-1-i] = current.data
// 			}

// 			if current.left != nil {
// 				queue = append(queue, current.left)
// 			}
// 			if current.right != nil {
// 				queue = append(queue, current.right)
// 			}
// 		}
// 		result = append(result, level)
// 		leftToRight = !leftToRight // swaitch direction for left to right
// 	}
// 	return result
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.left.right = &Node{data: 5}
// 	root.right.left = &Node{data: 6}
// 	root.right.right = &Node{data: 7}

// 	fmt.Println("Zigzag level order traversal : ")
// 	result := ZigzagLevelOrder(root)
// 	fmt.Println(result)
// }
