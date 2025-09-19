// Symmetric tree

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func Mirror(root *Node) bool {
	if root == nil {
		return true
	}

	queue := []*Node{root.left, root.right}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.data != right.data {
			return false
		}
		queue = append(queue, left.left)
		queue = append(queue, right.right)
		queue = append(queue, left.right)
		queue = append(queue, right.left)
	}
	return true
}

func main() {
	root := &Node{
		data: 1,
		left: &Node{
			data:  2,
			left:  &Node{data: 3},
			right: &Node{data: 4},
		},
		right: &Node{
			data:  2,
			left:  &Node{data: 4},
			right: &Node{data: 3},
		},
	}

	fmt.Println(Mirror(root))
}
