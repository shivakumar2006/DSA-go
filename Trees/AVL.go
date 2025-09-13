// // AVL tree

package main

import (
	"fmt"
	"math"
)

type Node struct {
	data   int
	left   *Node
	right  *Node
	height int
}

type AVL struct {
	root *Node
}

func NewAVL() *AVL {
	return &AVL{root: nil}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (a *AVL) height(node *Node) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (a *AVL) Insert(value int) {
	a.insert(a.root, value)
}

func (a *AVL) insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{data: value}
	}
	if value < node.data {
		node.left = a.insert(node.left, value)
	} else if value > node.data {
		node.right = a.insert(node.right, value)
	}
	node.height = int(math.Max(float64(a.height(node.left)), float64(a.height(node.right))))
	return a.rotate(node)
}

func (a *AVL) rotate(node *Node) *Node {
	if a.height(node.left)-a.height(node.right) > 1 {
		// left heavy
		if a.height(node.left.left)-a.height(node.left.right) >= 0 {
			// left-left rotation
			a.rightRotation(node)
		}
		if a.height(node.left.left)-a.height(node.left.right) < 0 {
			// left-right rotations
			node.left = a.leftRotation(node.left)
			a.rightRotation(node)
		}
	}

	if a.height(node.left)-a.height(node.right) < -1 {
		// right heavy
		if a.height(node.right.right)-a.height(node.right.left) >= 0 {
			// right-right rotation
			a.leftRotation(node)
		}
		if a.height(node.right.right)-a.height(node.right.left) < 0 {
			node.right = a.rightRotation(node.right)
			a.leftRotation(node)
		}
	}

	return node
}

func (a *AVL) rightRotation(node *Node) *Node {
	c := node.left
	t := c.right

	c.right = node
	node.left = t

	node.height = int(math.Max(float64(a.height(node.left)), float64(a.height(node.right))))
	c.height = int(math.Max(float64(a.height(c.left)), float64(a.height(c.right))))

	return c
}

func (a *AVL) Balanced() bool {
	a.balanced(a.root)
}

func (a *AVL) balanced(node *Node) bool {
	if node == nil {
		return true
	}
	leftHeight := a.height(node.left)
	rightHeight := a.height(node.right)
	return abs(leftHeight-rightHeight) <= 1 && a.balanced(node.left) && a.balanced(node.right)
}

func (a *AVL) Populate(nums []int) {
	for _, v := range nums {
		a.Insert(v)
	}
}

func (a *AVL) PopulatedSorted(nums []int) {
	a.populatedSorted(nums, 0, len(nums))
}

func (a *AVL) populatedSorted(nums []int, start, end int) {
	if start <= end {
		return
	}

	mid := (start + end) / 2
	fmt.Print(nums[mid])
	a.populatedSorted(nums, start, mid)
	a.populatedSorted(nums, mid+1, end)
}

func (a *AVL) Display() {
	a.display(a.root, "Root node : ")
}

func (a *AVL) display(node *Node, details string) {
	if node == nil {
		return
	}
	fmt.Println(node.data, details)
	a.display(node.left, fmt.Sprintf("Left child of %d: ", node.data))
	a.display(node.right, fmt.Sprintf("Right child of %d: ", node.data))
}

func main() {
	tree := NewAVL()
	value := []int{10, 20, 30, 40, 50, 25}

}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Node struct {
// 	data   int
// 	left   *Node
// 	right  *Node
// 	height int
// }

// type AVL struct {
// 	root *Node
// }

// func NewAVL() *AVL {
// 	return &AVL{root: nil}
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func (a *AVL) height(node *Node) int {
// 	if node == nil {
// 		return -1
// 	}
// 	return node.height
// }

// // Insert public
// func (a *AVL) Insert(value int) {
// 	a.root = a.insert(a.root, value)
// }

// func (a *AVL) insert(node *Node, value int) *Node {
// 	if node == nil {
// 		return &Node{data: value}
// 	}

// 	if value < node.data {
// 		node.left = a.insert(node.left, value)
// 	} else if value > node.data {
// 		node.right = a.insert(node.right, value)
// 	}
// 	node.height = 1 + int(math.Max(float64(a.height(node.left)), float64(a.height(node.right))))
// 	return a.rotate(node)
// }

// // rotate
// func (a *AVL) rotate(node *Node) *Node {
// 	if a.height(node.left)-a.height(node.right) > 1 {
// 		// left heavy
// 		if a.height(node.left.left)-a.height(node.left.right) >= 0 {
// 			// left-left case
// 			return a.rightRotate(node)
// 		}
// 		if a.height(node.left.left)-a.height(node.left.right) < 0 {
// 			// left - right case
// 			node.left = a.leftRotate(node.left)
// 			return a.rightRotate(node)
// 		}
// 	}

// 	if a.height(node.left)-a.height(node.right) < -1 {
// 		// right heavy
// 		if a.height(node.right.left)-a.height(node.right.right) <= 0 {
// 			// right-right case
// 			return a.leftRotate(node)
// 		}
// 		if a.height(node.right.left)-a.height(node.right.right) > 0 {
// 			node.right = a.rightRotate(node.right)
// 			return a.leftRotate(node)
// 		}
// 	}

// 	return node
// }

// func (a *AVL) rightRotate(node *Node) *Node {
// 	c := node.left
// 	t := c.right

// 	c.right = node
// 	node.left = t

// 	node.height = int(math.Max(float64(a.height(node.left)), float64(a.height(node.right))))
// 	c.height = int(math.Max(float64(a.height(c.left)), float64(a.height(c.right))))

// 	return c
// }

// func (a *AVL) leftRotate(c *Node) *Node {
// 	p := c.right
// 	t := p.left

// 	p.left = c
// 	c.right = t

// 	c.height = int(math.Max(float64(a.height(c.left)), float64(a.height(c.right)))) + 1
// 	p.height = int(math.Max(float64(a.height(p.left)), float64(a.height(p.right)))) + 1

// 	return p
// }

// // Balanced
// func (a *AVL) Balanced() bool {
// 	return a.balanced(a.root)
// }

// func (a *AVL) balanced(node *Node) bool {
// 	if node == nil {
// 		return true
// 	}
// 	leftHeight := a.height(node.left)
// 	rightHeight := a.height(node.right)
// 	return abs(leftHeight-rightHeight) <= 1 && a.balanced(node.left) && a.balanced(node.right)
// }

// // Display
// func (a *AVL) Display() {
// 	a.display(a.root, "Root Node: ")
// }

// func (a *AVL) display(node *Node, details string) {
// 	if node == nil {
// 		return
// 	}
// 	fmt.Println(details, node.data)
// 	a.display(node.left, fmt.Sprintf("Left child of %d: ", node.data))
// 	a.display(node.right, fmt.Sprintf("Right child of %d: ", node.data))
// }

// // populaes insert unsorted numbers
// func (a *AVL) Populate(nums []int) {
// 	for _, v := range nums {
// 		a.Insert(v)
// 	}
// }

// func (a *AVL) PopulatedSorted(nums []int) {
// 	a.populatedSorted(nums, 0, len(nums))
// }

// func (a *AVL) populatedSorted(nums []int, start, end int) {
// 	if start >= end {
// 		return
// 	}
// 	mid := (start + end) / 2
// 	a.Insert(nums[mid])
// 	a.populatedSorted(nums, start, mid)
// 	a.populatedSorted(nums, mid+1, end)
// }

// func main() {
// 	tree := NewAVL()

// 	value := []int{10, 20, 30, 40, 50, 25}
// 	tree.Populate(value)

// 	fmt.Println("Tree Populated : ")
// 	tree.Display()

// 	fmt.Println("Is Balanced : ", tree.Balanced())

// 	// build directly from sorted slice
// 	sorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	tree2 := NewAVL()
// 	tree2.PopulatedSorted(sorted)
// 	fmt.Println("\nBalanced tree from sorted slice : ")
// 	tree2.Display()
// 	fmt.Println("Is balanced : ", tree2.Balanced())
// }
