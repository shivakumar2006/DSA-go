// Implementation of Binary Tree

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func readInt(reader *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	b, _ := strconv.Atoi(text)
	return b
}

func readBool(reader *bufio.Reader, prompt string) bool {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	num, _ := strconv.ParseBool(text)
	return num
}

func (t *BinaryTree) populate(reader *bufio.Reader, node *Node) {
	left := readBool(reader, fmt.Sprintf("Do you want to enter left of %d: ", node.data))
	if left {
		value := readInt(reader, fmt.Sprintf("Enter the value in left of %d: ", node.data))
		node.left = &Node{data: value}
		t.populate(reader, node.left)
	}

	right := readBool(reader, fmt.Sprintf("Do you want to enter right of %d: ", node.data))
	if right {
		value := readInt(reader, fmt.Sprintf("Enter the value in right of %d: ", node.data))
		node.right = &Node{data: value}
		t.populate(reader, node.right)
	}
}

func (t *BinaryTree) prettyDisplay(node *Node, level int) {
	if node == nil {
		return
	}

	t.prettyDisplay(node.right, level+1)

	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|\t\t")
		}
		fmt.Println("|------->", node.data)
	} else {
		fmt.Println(node.data)
	}
	t.prettyDisplay(node.left, level+1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	tree := &BinaryTree{}

	rootValue := readInt(reader, "Enter the root value : ")
	tree.root = &Node{data: rootValue}

	tree.populate(reader, tree.root)

	fmt.Println("\nPretty Display : ")
	tree.prettyDisplay(tree.root, 0)
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// type Node struct {
// 	data  int
// 	left  *Node
// 	right *Node
// }

// type BinaryTrees struct {
// 	root *Node
// }

// // Helper to read int input
// func readBool(reader *bufio.Reader, prompt string) bool {
// 	fmt.Print(prompt)
// 	text, _ := reader.ReadString('\n')
// 	text = strings.TrimSpace(text)
// 	b, _ := strconv.ParseBool(text)
// 	return b
// }

// // helper to read int input
// func readInt(reader *bufio.Reader, prompt string) int {
// 	fmt.Print(prompt)
// 	text, _ := reader.ReadString('\n')
// 	text = strings.TrimSpace(text)
// 	num, _ := strconv.Atoi(text)
// 	return num
// }

// // Populate starts for building of the tree
// func (t *BinaryTrees) populate(reader *bufio.Reader, node *Node) {
// 	left := readBool(reader, fmt.Sprintf("Do you want to enter left of %d (true/false): ", node.data))
// 	if left {
// 		value := readInt(reader, fmt.Sprintf("Enter the value of the left of %d: ", node.data))
// 		node.left = &Node{data: value}
// 		t.populate(reader, node.left)
// 	}

// 	right := readBool(reader, fmt.Sprintf("Do you want to enter right of %d (true/false): ", node.data))
// 	if right {
// 		value := readInt(reader, fmt.Sprintf("Enter the value of the right of %d: ", node.data))
// 		node.right = &Node{data: value}
// 		t.populate(reader, node.right)
// 	}
// }

// // Display prints tree value with indentations
// func (t *BinaryTrees) Display() {
// 	t.display(t.root, "")
// }

// func (t *BinaryTrees) display(node *Node, indent string) {
// 	if node == nil {
// 		return
// 	}
// 	fmt.Println(indent + strconv.Itoa(node.data))
// 	t.display(node.left, indent+"\t")
// 	t.display(node.right, indent+"\t")
// }

// func (t *BinaryTrees) preetyDisplay(node *Node, level int) {
// 	if node == nil {
// 		return
// 	}

// 	t.preetyDisplay(node.right, level+1)

// 	if level != 0 {
// 		for i := 0; i < level-1; i++ {
// 			fmt.Print("|\t\t")
// 		}
// 		fmt.Println("|--------->", node.data)
// 	} else {
// 		fmt.Println(node.data)
// 	}
// 	t.preetyDisplay(node.left, level+1)
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	tree := &BinaryTrees{}

// 	rootValue := readInt(reader, "Enter the root value : ")
// 	tree.root = &Node{data: rootValue}

// 	tree.populate(reader, tree.root)

// 	fmt.Println("\nIndented Display : ")
// 	tree.Display()

// 	fmt.Println("\nPretty Display : ")
// 	tree.preetyDisplay(tree.root, 0)
// }
