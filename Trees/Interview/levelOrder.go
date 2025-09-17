// // Level order successor of a node...
// // Google question

package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func levelOrderSuccessor(root *Node, key int) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}

		if current.data == key {
			if len(queue) > 0 {
				return queue[0]
			}
			return nil
		}
	}
	return nil
}

func main() {
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}
	root.right.left = &Node{data: 6}
	root.right.right = &Node{data: 7}

	key := 3
	successor := levelOrderSuccessor(root, key)
	if successor != nil {
		fmt.Println("Level order of", key, "is", successor.data)
	} else {
		fmt.Println("No successor is found for", key)
	}

}

// package main

// import "fmt"

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// func levelOrderSuccessor(root *Node, key int) *Node {
// 	if root == nil {
// 		return nil
// 	}

// 	queue := []*Node{root}

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]

// 		if current.left != nil {
// 			queue = append(queue, current.left)
// 		}
// 		if current.right != nil {
// 			queue = append(queue, current.right)
// 		}

// 		// if this is the key the next node in queue is the successor
// 		if current.data == key {
// 			if len(queue) > 0 {
// 				return queue[0]
// 			}
// 			return nil // no successor
// 		}
// 	}
// 	return nil
// }

// func main() {
// 	root := &Node{data: 1}
// 	root.left = &Node{data: 2}
// 	root.right = &Node{data: 3}
// 	root.left.left = &Node{data: 4}
// 	root.left.right = &Node{data: 5}
// 	root.right.left = &Node{data: 6}
// 	root.right.right = &Node{data: 7}

// 	key := 3
// 	successor := levelOrderSuccessor(root, key)
// 	if successor != nil {
// 		fmt.Println("Level order successor of", key, "is", successor.data)
// 	} else {
// 		fmt.Println("No successor found for", key)
// 	}
// }
