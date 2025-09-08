package main

import "fmt"

func main() {
	queue := []int{10, 20, 30, 40, 50}

	fmt.Println("Queue : ", queue)

	// remove one element from front
	removed := queue[0]
	queue = queue[1:]

	fmt.Println("Removed element : ", removed)
	fmt.Println("Queue after removal. : ", queue)
}
