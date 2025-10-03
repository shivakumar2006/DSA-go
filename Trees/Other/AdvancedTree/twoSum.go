// Two Sum 4
// Google

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func findTarget(root *Node, k int) bool {
	seen := make(map[int]bool)
	return helper(root, k, seen)
}

func helper(root *Node, k int, seen map[int]bool) bool {
	if root == nil {
		return false
	}
	if seen[k-root.data] {
		return true
	}

	seen[root.data] = true
	return helper(root.left, k, seen) || helper(root.right, k, seen)
}

func main() {
	root := &Node{
		data: 5,
		left: &Node{
			data:  3,
			left:  &Node{data: 2},
			right: &Node{data: 4},
		},
		right: &Node{
			data:  6,
			right: &Node{data: 7},
		},
	}

	fmt.Println(findTarget(root, 9))  // true (2 + 7)
	fmt.Println(findTarget(root, 28)) // false
}
