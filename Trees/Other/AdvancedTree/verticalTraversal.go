// vertical order traversal of a bianary Tree
// Google interview question

package main

import (
	"fmt"
	"sort"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func verticalTraversal(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	colMaps := make(map[int][]int)
	minCol, maxCol := 0, 0

	type QueueNode struct {
		node *Node
		col  int
	}
	queue := []QueueNode{{node: root, col: 0}}

	for len(queue) > 0 {
		nextQueue := []QueueNode{}
		for _, qn := range queue {
			node, col := qn.node, qn.col
			colMaps[col] = append(colMaps[col], node.val)
			if node.left != nil {
				nextQueue = append(nextQueue, QueueNode{node.left, col - 1})
			}
			if node.right != nil {
				nextQueue = append(nextQueue, QueueNode{node.right, col + 1})
			}

			if col < minCol {
				minCol = col
			}
			if col > maxCol {
				maxCol = col
			}
		}
		queue = nextQueue
	}
	ans := [][]int{}
	for c := minCol; c <= maxCol; c++ {
		sort.Ints(colMaps[c])
		ans = append(ans, colMaps[c])
	}
	return ans
}

func main() {
	root := &Node{val: 3}
	root.left = &Node{val: 9}
	root.right = &Node{val: 20}
	root.right.left = &Node{val: 15}
	root.right.right = &Node{val: 7}

	result := verticalTraversal(root)
	fmt.Println(result) // Output: [[9], [3,15], [20], [7]]
}
