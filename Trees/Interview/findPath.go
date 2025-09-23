// path exist in binarry tree from root to left

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func findPath(root *Node, arr []int) bool {
	if root == nil {
		return len(arr) == 0
	}
	return findPathHelper(root, arr, 0)
}

func findPathHelper(root *Node, arr []int, index int) bool {
	if root == nil {
		return false
	}

	if index >= len(arr) || root.data != arr[index] {
		return false
	}

	if root.left == nil && root.right == nil && index == len(arr)-1 {
		return true
	}

	return findPathHelper(root.left, arr, index+1) || findPathHelper(root.right, arr, index+1)
}

// Count all paths whose sum equals a given value
func countPaths(root *Node, sum int) int {
	var path []int
	return countPathsHelper(root, sum, path)
}

func countPathsHelper(root *Node, target int, path []int) int {
	if root == nil {
		return 0
	}

	path = append(path, root.data)

	count := 0
	sum := 0

	// walk backwards through path to check sub-paths sums
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == target {
			count++
		}
	}

	count += countPathsHelper(root.left, target, path)
	count += countPathsHelper(root.right, target, path)

	// backtrack manually
	path = path[:len(path)-1]

	return count
}

func display(root *Node, lev int) {
	if root == nil {
		return
	}

	display(root.right, lev+1)
	if lev != 0 {
		for i := 0; i < lev-1; i++ {
			fmt.Print("|\t")
		}
		fmt.Println("|------>", root.data)
	} else {
		fmt.Println(root.data)
	}
	display(root.left, lev+1)
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

	arr := []int{5, 4, 11, 2}
	fmt.Println("path exists:", findPath(root, arr)) // true
	fmt.Println()

	fmt.Println("count of paths with sum 22:", countPaths(root, 22))
	fmt.Println()

	fmt.Println("Display tree : ")
	display(root, 0)
}

// func findPath(root *Node, arr []int) bool {
// 	if root == nil {
// 		return len(arr) == 0
// 	}
// 	return findPathHelper(root, arr, 0)
// }

// func findPathHelper(root *Node, arr []int, index int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if index >= len(arr) || root.data != arr[index] {
// 		return false
// 	}

// 	if root.left == nil && root.right == nil && index == len(arr)-1 {
// 		return true
// 	}

// 	return findPathHelper(root.left, arr, index+1) || findPathHelper(root.right, arr, index+1)
// }

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// // return true if there is a root-to-leaf path that matches array
// func findPath(root *Node, arr []int) bool {
// 	if root == nil {
// 		return len(arr) == 0
// 	}
// 	return findPathHelper(root, arr, 0)
// }

// func findPathHelper(root *Node, arr []int, index int) bool {
// 	if root == nil {
// 		return false
// 	}

// 	if index >= len(arr) || root.data != arr[index] {
// 		return false
// 	}

// 	// if leaf node at last index of array
// 	if root.left == nil && root.right == nil && index == len(arr)-1 {
// 		return true
// 	}

// 	return findPathHelper(root.left, arr, index+1) || findPathHelper(root.right, arr, index+1)
// }
