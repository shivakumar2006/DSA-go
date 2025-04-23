package main

import (
	"fmt"
)

func main() {
	arr := []int{-2, 5, -3, 7, 69, -100, 45}

	findMin := min(arr)
	fmt.Printf("The minimum number in the array is %d\n", findMin)
}

// func min(arr []int) int {
// 	if len(arr) == 0 {
// 		panic("array is empty")
// 	}

// 	smallest := arr[0]
// 	for _, value := range arr {
// 		if value < smallest {
// 			smallest = value
// 		}
// 	}
// 	return smallest
// }

func min(arr []int) int {
	if len(arr) == 0 {
		panic("array is empty")
	}

	smallest := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}
	return smallest
}
