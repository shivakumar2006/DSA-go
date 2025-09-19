// Binary tree right side view
// Amazon & Flipkart question

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func RightSideView(root *Node) []int {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	result := []int{}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				result = append(result, current.data)
			}

			if current.left != nil {
				queue = append(queue, current.left)
			}
			if current.right != nil {
				queue = append(queue, current.right)
			}
		}
	}
	return result
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.right = &Node{data: 6}

	fmt.Println(RightSideView(root)) // [1 3 6]
}
