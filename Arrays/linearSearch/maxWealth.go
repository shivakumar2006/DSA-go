package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{1, 2, 3},
		{3, 4, 5},
		{2, 2, 2},
	}

	res := maxWealth(arr)
	fmt.Printf("The maximum wealth is : %d\n", res)
}

func maxWealth(arr [][]int) int {
	max := 0
	for row := 0; row < len(arr); row++ {
		sum := 0
		for col := 0; col < len(arr[row]); col++ {
			sum += arr[row][col]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [][]int{
// 		{1, 2, 3},
// 		{3, 4, 5},
// 		{2, 2, 2},
// 	}
// 	res := maxWealth(arr)
// 	fmt.Printf("The maximum wealth is : %d\n", res)
// }

// func maxWealth(arr [][]int) int {
// 	max := 0
// 	for row := 0; row < len(arr); row++ {
// 		sum := 0
// 		for col := 0; col < len(arr[row]); col++ {
// 			sum += arr[row][col]
// 		}
// 		if sum > max {
// 			max = sum
// 		}
// 	}
// 	return max
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [][]int{
// 		{1, 2, 3},
// 		{3, 4, 5},
// 		{2, 2, 2},
// 	}

// 	res := maxWealth(arr)
// 	fmt.Printf("The maximum wealth is  : %d\n", res)
// }

// func maxWealth(arr [][]int) int {
// 	max := 0
// 	for row := 0; row < len(arr); row++ {
// 		sum := 0
// 		for col := 0; col < len(arr[row]); col++ {
// 			sum += arr[row][col]
// 		}
// 		if sum > max {
// 			max = sum
// 		}
// 	}
// 	return max
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := [][]int{
// 		{1, 2, 3}, // total: 6
// 		{3, 4, 5}, // total: 12
// 		{2, 2, 2}, // total: 6
// 	}

// 	result := maxWealth(arr)
// 	fmt.Printf("The maximum wealth is  : %d\n", result)
// }

// func maxWealth(arr [][]int) int {
// 	max := 0
// 	for row := 0; row < len(arr); row++ {
// 		sum := 0
// 		for col := 0; col < len(arr[row]); col++ {
// 			sum += arr[row][col]
// 		}
// 		if sum > max {
// 			max = sum
// 		}
// 	}
// 	return max
// }
