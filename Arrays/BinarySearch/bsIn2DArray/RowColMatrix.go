// find the target in 2d array(Matrix)

package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{10, 20, 30, 40},
		{15, 25, 35, 45},
		{28, 29, 37, 49},
		{33, 34, 38, 50},
	}
	tar := 49
	res := search(arr, tar)
	if res[0] != -1 {
		fmt.Printf("The target is found at the index : [%d, %d]\n", res[0], res[1])
	} else {
		fmt.Println("Target not found")
	}
}

func search(arr [][]int, tar int) []int {
	row := 0
	col := len(arr) - 1
	for row < len(arr) && col >= 0 {
		if arr[row][col] == tar {
			return []int{row, col}
		} else if arr[row][col] < tar {
			row++
		} else {
			col--
		}
	}
	return []int{-1, -1}
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [][]int{
// 		{10, 20, 30, 40},
// 		{15, 25, 35, 45},
// 		{28, 29, 37, 49},
// 		{33, 34, 38, 50},
// 	}
// 	tar := 49
// 	res := search(arr, tar)
// 	if res[0] != -1 {
// 		fmt.Printf("Target is found at the index : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not foound")
// 	}
// }

// func search(arr [][]int, tar int) []int {
// 	row := 0
// 	col := len(arr) - 1
// 	for row < len(arr) && col >= 0 {
// 		if arr[row][col] == tar {
// 			return []int{row, col}
// 		} else if arr[row][col] < tar {
// 			row++
// 		} else {
// 			col--
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
// 		{10, 20, 30, 40},
// 		{15, 25, 35, 45},
// 		{28, 29, 37, 49},
// 		{33, 34, 38, 50},
// 	}
// 	tar := 29
// 	res := search(arr, tar)
// 	if res[0] != -1 {
// 		fmt.Printf("The target is found at the index : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr [][]int, tar int) []int {
// 	row := 0
// 	col := len(arr) - 1
// 	for row < len(arr) && col >= 0 {
// 		if arr[row][col] == tar {
// 			return []int{row, col}
// 		} else if arr[row][col] < tar {
// 			row++
// 		} else {
// 			col--
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
// 		{10, 20, 30, 40},
// 		{15, 25, 35, 45},
// 		{28, 29, 37, 49},
// 		{33, 34, 38, 50},
// 	}
// 	tar := 38
// 	res := search(arr, tar)
// 	if res[0] != -1 {
// 		fmt.Printf("Target is found : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr [][]int, tar int) []int {
// 	row := 0
// 	col := len(arr) - 1
// 	for row < len(arr) && col >= 0 {
// 		if arr[row][col] == tar {
// 			return []int{row, col}
// 		}
// 		if arr[row][col] < tar {
// 			row++
// 		} else {
// 			col--
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
// 		{10, 20, 30, 40},
// 		{15, 25, 35, 45},
// 		{28, 29, 37, 49},
// 		{33, 34, 38, 50},
// 	}
// 	tar := 38
// 	res := search(arr, tar)
// 	if res[0] != -1 {
// 		fmt.Printf("Target is found : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target nor found")
// 	}
// }

// func search(arr [][]int, tar int) []int {
// 	row := 0
// 	col := len(arr) - 1
// 	for row < len(arr) && col >= 0 {
// 		if arr[row][col] == tar {
// 			return []int{row, col}
// 		}
// 		if arr[row][col] < tar {
// 			row++
// 		} else {
// 			col--
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
// 		{10, 20, 30, 40},
// 		{15, 25, 35, 45},
// 		{28, 29, 37, 49},
// 		{33, 34, 38, 50},
// 	}
// 	tar := 38
// 	res := search(arr, tar)
// 	if res[0] != -1 {
// 		fmt.Printf("The target is found : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr [][]int, tar int) []int {
// 	row := 0
// 	col := len(arr) - 1
// 	for row < len(arr) && col >= 0 {
// 		if arr[row][col] == tar {
// 			return []int{row, col}
// 		}
// 		if arr[row][col] < tar {
// 			row++
// 		} else {
// 			col--
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
// 		{10, 20, 30, 40},
// 		{15, 25, 35, 45},
// 		{28, 29, 37, 49},
// 		{33, 34, 38, 50},
// 	}
// 	tar := 37
// 	res := search(arr, tar)
// 	if res[0] != -1 {
// 		fmt.Printf("The target found : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr [][]int, tar int) []int {
// 	row := 0
// 	col := len(arr[0]) - 1
// 	for row < len(arr) && col >= 0 {
// 		if arr[row][col] == tar {
// 			return []int{row, col}
// 		}
// 		if arr[row][col] < tar {
// 			row++
// 		} else {
// 			col--
// 		}
// 	}
// 	return []int{-1, -1}
// }
