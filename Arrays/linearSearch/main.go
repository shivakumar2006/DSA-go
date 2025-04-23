package main

import (
	"fmt"
)

func main() {
	// linearSearch on te array
	arr := []int{-2, 5, -3, 7, 69, -100, 45}
	target := -100

	result := linearSearch(arr, target)
	if result != -1 {
		fmt.Printf("Element is found at the index of %d\n", result)
	} else {
		fmt.Println("Element is not found")
	}

	// linearSearch on the string
	str := "shiva"
	targetChar := "v"

	resultStr := search(str, targetChar)
	if result != -1 {
		fmt.Printf("Element is found at the index of %d\n", resultStr)
	} else {
		fmt.Println("Element is not found")
	}
}
