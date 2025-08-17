package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 1, 4, 3, 2, 6}
	res := sort(arr)
	fmt.Println(res)
}

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				swap(arr, j, j-1)
			} else {
				break
			}
		}
	}
	return arr
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := i + 1; j > 0; j-- {
// 			if arr[j] < arr[j-1] {
// 				swap(arr, j, j-1)
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := i + 1; j > 0; j-- {
// 			if arr[j] < arr[j-1] {
// 				swap(arr, j, j-1)
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := i + 1; j > 0; j-- {
// 			if arr[j] < arr[j-1] {
// 				swap(arr, j, j-1)
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, j int, i int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := i + 1; j > 0; j-- {
// 			if arr[j] < arr[j-1] {
// 				swap(arr, j, j-1)
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, j int, i int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := i + 1; j > 0; j-- {
// 			if arr[j] < arr[j-1] {
// 				swap(arr, j, j-1)
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, j int, i int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }
