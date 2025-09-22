// Searialize and Desearialize Binary tree
// Google, Amazon and Facebook interview

package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func serialize(root *Node) []string {
	var list []string
	serializeHelper(root, &list)
	return list
}

func serializeHelper(node *Node, list *[]string) {
	if node == nil {
		*list = append(*list, "null")
		return
	}

	*list = append(*list, strconv.Itoa(node.data))
	serializeHelper(node.left, list)
	serializeHelper(node.right, list)
}

func deserialize(list []string) *Node {
	// reverse the list in-place
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return deseriallizeHelper(&list)
}

func deseriallizeHelper(list *[]string) *Node {
	// pop from the end ( since we reversed )
	if len(*list) == 0 {
		return nil
	}

	val := (*list)[len(*list)-1]
	*list = (*list)[:len(*list)-1]

	if val == "null" {
		return nil
	}

	num, _ := strconv.Atoi(val)
	node := &Node{data: num}
	node.left = deseriallizeHelper(list)
	node.right = deseriallizeHelper(list)
	return node
}

func preorder(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.data, " ")
	preorder(root.left)
	preorder(root.right)
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
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.right.left = &Node{data: 4}
	root.right.right = &Node{data: 5}

	// serialize
	list := serialize(root)
	fmt.Println("Serialized:", list)

	// deserialize
	newRoot := deserialize(list)

	fmt.Print("Preorder of Deserialized Tree: ")
	preorder(newRoot)
	fmt.Println()

	fmt.Println("Display tree after deserialized: ")
	display(root, 0)
	fmt.Println()
}
