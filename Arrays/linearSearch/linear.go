package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 5, 3, 7, 69, 100, 45}
	target := 69

	result := linearSearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found")
	}
}

func linearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	return -1
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	array := []int{1, 4, 5, 6, 3, 7, 10, 69}
// 	target := 7

// 	result := linearSearch(array, target)

// 	if result != -1 {
// 		fmt.Printf("Element found at index %d\n", result)
// 	} else {
// 		fmt.Println("Element not found")
// 	}
// }

// func linearSearch(array []int, target int) int {
// 	for i, value := range array {
// 		if value == target {
// 			return i
// 		}
// 	}
// 	return -1
// }
