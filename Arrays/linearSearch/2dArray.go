// //linearSearch in 2D Array...

package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{23, 4, 1},
		{57, 69, 29},
		{48, 99, 72, 55},
		{20, 4, 54},
	}
	target := 55
	res := search(arr, target)

	if res != nil {
		fmt.Printf("Element is found at the index of [%d][%d]\n", res[0], res[1])
	} else {
		fmt.Println("Element is not found")
	}
}

func search(arr [][]int, target int) []int {
	for row := 0; row < len(arr); row++ {
		for col := 0; col < len(arr[row]); col++ {
			if arr[row][col] == target {
				return []int{row, col}
			}
		}
	}
	return nil
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [][]int{
// 		{23, 4, 1},
// 		{57, 69, 29},
// 		{48, 99, 72, 55},
// 		{20, 4, 54},
// 	}

// 	target := 55
// 	res := search(arr, target)

// 	if res != nil {
// 		fmt.Printf("Element is found at the index of [%d][%d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Element is not found")
// 	}
// }

// func search(arr [][]int, target int) []int {
// 	for row := 0; row < len(arr); row++ {
// 		for col := 0; col < len(arr[row]); col++ {
// 			if arr[row][col] == target {
// 				return []int{row, col}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [][]int{
// 		{23, 4, 1},
// 		{57, 69, 29},
// 		{48, 99, 72, 55},
// 		{20, 4, 54},
// 	}

// 	target := 69
// 	result := search(arr, target)

// 	if result != nil {
// 		fmt.Printf("Element is found at the index of [%d][%d]\n", result[0], result[1])
// 	} else {
// 		fmt.Println("Element is not found")
// 	}
// }

// func search(arr [][]int, target int) []int {
// 	for row := 0; row < len(arr); row++ {
// 		for col := 0; col < len(arr[row]); col++ {
// 			if arr[row][col] == target {
// 				return []int{row, col}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [][]int{
// 		{23, 4, 1},
// 		{57, 69, 29},
// 		{48, 99, 72, 55},
// 		{20, 4, 54},
// 	}
// 	target := 72
// 	result := search(arr, target)
// 	fmt.Printf("Element is found at the index of [%d][%d]\n", result[0], result[1])
// }

// func search(arr [][]int, target int) []int {
// 	for row := 0; row < len(arr); row++ {
// 		for col := 0; col < len(arr[row]); col++ {
// 			if arr[row][col] == target {
// 				return []int{row, col}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }
