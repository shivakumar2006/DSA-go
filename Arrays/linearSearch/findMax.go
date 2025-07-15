package main

import "fmt"

func main() {
	arr := []int{-2, 5, 7, 69, -100}
	findMax := max(arr)
	fmt.Printf("THe max no in the array is : %d\n", findMax)
}

func max(arr []int) int {
	if len(arr) == 0 {
		panic("array is empty")
	}

	lar := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > lar {
			lar = arr[i]
		}
	}
	return lar
}

// package main

// import "fmt"

// func main() {
// 	arr := []int{-2, 5, 7, 69, -100, 45}
// 	findMax := max(arr)
// 	fmt.Printf("the max no in the array is : %d\n", findMax)
// }

// func max(arr []int) int {
// 	if len(arr) == 0 {
// 		panic("array is empty")
// 	}

// 	lar := arr[0]
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] > lar {
// 			lar = arr[i]
// 		}
// 	}
// 	return lar
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{-2, 5, -3, 7, 69, -100, 45}

// 	findMax := max(arr)
// 	fmt.Printf("The maximum number is %d\n", findMax)
// }

// func max(arr []int) int {
// 	if len(arr) == 0 {
// 		panic("array is empty")
// 	}

// 	largest := arr[0]
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] > largest {
// 			largest = arr[i]
// 		}
// 	}
// 	return largest
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{-2, 5, -3, 7, 69, -100, 45}

// 	findMax := max(arr)
// 	fmt.Printf("The maximum number is %d\n", findMax)
// }

// func max(arr []int) int {
// 	if len(arr) == 0 {
// 		panic("Array is empty")
// 	}

// 	largest := arr[0]
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] > largest {
// 			largest = arr[i]
// 		}
// 	}
// 	return largest
// }

// func max(arr []int) int {
// 	if len(arr) == 0 {
// 		panic("Array is empty")
// 	}

// 	largest := arr[0]
// 	for _, value := range arr {
// 		if value > largest {
// 			largest = value
// 		}
// 	}
// 	return largest
// }
