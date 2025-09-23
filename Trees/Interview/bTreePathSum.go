// Binary tree path sum
// Facebook question \

package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

var ans int

func maxPathSum(root *Node) int {
	ans = math.MinInt
	helper(root)
	return ans
}

func helper(root *Node) int {
	if root == nil {
		return 0
	}

	left := helper(root.left)
	right := helper(root.right)

	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}

	pathSum := left + right + root.data

	if pathSum > ans {
		ans = pathSum
	}

	if left > right {
		return left + root.data
	} else {
		return right + root.data
	}
}

func display(root *Node, level int) {
	if root == nil {
		return
	}

	display(root.right, level+1)
	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|\t")
		}
		fmt.Println("|------>", root.data)
	} else {
		fmt.Println(root.data)
	}
	display(root.left, level+1)
}

func main() {
	root := &Node{data: -10}
	root.left = &Node{data: 9}
	root.right = &Node{data: 20}
	root.right.left = &Node{data: 15}
	root.right.right = &Node{data: 7}

	fmt.Println("max path sum : ", maxPathSum(root))
	fmt.Println()

	fmt.Println("Display tree : ")
	display(root, 0)
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// var ans int

// func maxPathSum(root *Node) int {
// 	ans = math.MinInt
// 	helper(root)
// 	return ans
// }

// func helper(root *Node) int {
// 	if root == nil {
// 		return 0
// 	}

// 	left := helper(root.left)
// 	right := helper(root.right)

// 	if left < 0 {
// 		return left
// 	}
// 	if right < 0 {
// 		return right
// 	}

// 	pathSum := left + right + root.data

// 	if pathSum > ans {
// 		ans = pathSum
// 	}

// 	if left > right {
// 		return left + root.data
// 	} else {
// 		return right + root.data
// 	}
// }

// func display(root *Node, level int) {
// 	if root == nil {
// 		return
// 	}

// 	display(root.right, level+1)
// 	if level != 0 {
// 		for i := 0; i < level-1; i++ {
// 			fmt.Print("|\t")
// 		}
// 		fmt.Println("|----->", root.data)
// 	} else {
// 		fmt.Println(root.data)
// 	}
// 	display(root.left, level+1)
// }

// func main() {
// 	root := &Node{data: -10}
// 	root.left = &Node{data: 9}
// 	root.right = &Node{data: 20}
// 	root.right.left = &Node{data: 15}
// 	root.right.right = &Node{data: 7}

// 	fmt.Println("Maximum Path Sum:", maxPathSum(root)) // Output: 42
// 	fmt.Println()

// 	fmt.Println("Display tree : ")
// 	display(root, 0)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// var ans int

// func maxPathSum(root *Node) int {
// 	ans = math.MinInt
// 	helper(root)
// 	return ans
// }

// func helper(root *Node) int {
// 	if root == nil {
// 		return 0
// 	}

// 	left := helper(root.left)
// 	right := helper(root.right)

// 	if left < 0 {
// 		return left
// 	}
// 	if right < 0 {
// 		return right
// 	}

// 	pathSum := left + right + root.data

// 	if pathSum > ans {
// 		ans = pathSum
// 	}

// 	if left > right {
// 		return left + root.data
// 	} else {
// 		return right + root.data
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// var ans int // global maximum

// func maxPathSum(root *Node) int {
// 	ans = math.MinInt // initialize to smallest unit
// 	helper(root)
// 	return ans
// }

// func helper(node *Node) int {
// 	if node == nil {
// 		return 0
// 	}

// 	// compute max path sum starting from left/right child
// 	left := helper(node.left)
// 	right := helper(node.right)

// 	// ignore negative paths
// 	if left < 0 {
// 		left = 0
// 	}
// 	if right < 0 {
// 		right = 0
// 	}

// 	// path sum going through this node
// 	pathSum := left + right + node.data

// 	// update global maximum
// 	if pathSum > ans {
// 		ans = pathSum
// 	}

// 	// return the max single path
// 	if left > right {
// 		return left + node.data
// 	} else {
// 		return right + node.data
// 	}
// }
