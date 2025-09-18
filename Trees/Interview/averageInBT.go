// // Average in Binary Tree

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func AverageOfLevels(root *Node) []float64 {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	var averages []float64

	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			sum += current.data

			if current.left != nil {
				queue = append(queue, current.left)
			}
			if current.right != nil {
				queue = append(queue, current.right)
			}
		}
		avg := float64(sum) / float64(levelSize)
		averages = append(averages, avg)
	}
	return averages
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}

	average := AverageOfLevels(root)
	fmt.Println("Average of each level", average)
}

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func AverageOfLevels(root *Node) []float64 {
// 	if root == nil {
// 		return nil
// 	}

// 	queue := []*Node{root}
// 	var averages []float64

// 	for len(queue) > 0 {
// 		levelSize := len(queue)
// 		sum := 0

// 		for i := 0; i < levelSize; i++ {
// 			current := queue[0]
// 			queue = queue[1:]

// 			sum += current.data

// 			if current.left != nil {
// 				queue = append(queue, current.left)
// 			}
// 			if current.right != nil {
// 				queue = append(queue, current.right)
// 			}
// 		}
// 		avg := float64(sum) / float64(levelSize)
// 		averages = append(averages, avg)
// 	}
// 	return averages
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.left.right = &Node{data: 5}
// 	root.right.left = &Node{data: 6}
// 	root.right.right = &Node{data: 7}

// 	Average := AverageOfLevels(root)
// 	fmt.Println("Average of each level: ", Average)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func AverageOfLevels(root *Node) []float64 {
// 	if root == nil {
// 		return nil
// 	}

// 	queue := []*Node{root}
// 	var averages []float64

// 	for len(queue) > 0 {
// 		levelSize := len(queue)
// 		sum := 0

// 		for i := 0; i < levelSize; i++ {
// 			current := queue[0]
// 			queue = queue[1:]

// 			sum += current.data

// 			if current.left != nil {
// 				queue = append(queue, current.left)
// 			}
// 			if current.right != nil {
// 				queue = append(queue, current.right)
// 			}
// 		}
// 		avg := float64(sum) / float64(levelSize)
// 		averages = append(averages, avg)
// 	}
// 	return averages
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.left.right = &Node{data: 5}
// 	root.right.left = &Node{data: 6}
// 	root.right.right = &Node{data: 7}

// 	averages := AverageOfLevels(root)
// 	fmt.Println("Average of each level : ", averages)
// }
