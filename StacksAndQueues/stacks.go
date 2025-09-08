// Stacks & Queues

package main

import "fmt"

func main() {
	stack := []int{}

	stack = append(stack, 10)
	stack = append(stack, 20)
	stack = append(stack, 30)
	stack = append(stack, 40)
	stack = append(stack, 50)

	fmt.Println("Stacks : ", stack)

	// pop all elements
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println("Popped : ", top)
	}

	fmt.Println("Stack after popping all the elements : ", stack)

	// // pop
	// top := stack[len(stack)-1]
	// stack = stack[:len(stack)-1]
	// fmt.Println("Popped : ", top)
	// fmt.Println("Stack now : ", stack)
}
