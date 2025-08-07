// First missing positive,
// Amazon

package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 2, 3, -1, 6}
	fmt.Println(sort(arr))
}

func sort(arr []int) int {
	i := 0
	for i < len(arr) {
		correct := arr[i] - 1
		if arr[i] > 0 && arr[i] <= len(arr) && arr[i] != arr[correct] {
			swap(arr, i, correct)
		} else {
			i++
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}
	return len(arr)
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// func sort(arr []int) int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] > 0 && arr[i] <= len(arr) && arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}

// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i+1 {
// 			return i + 1
// 		}
// 	}
// 	return len(arr)
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }
