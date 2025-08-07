package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 1, 4, 2, 5, 6}
	res := sort(arr)
	fmt.Println(res)
}

func sort(arr []int) []int {
	i := 0
	for i < len(arr) {
		correct := arr[i] - 1
		if arr[i] != arr[correct] {
			swap(arr, i, correct)
		} else {
			i++
		}
		// [1, 2, 2, 4, 5, 6] // 2 is duplicate and 3 is missing
	}
	var result []int
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			result = append(result, arr[i], i+1)
		}
	}
	return result
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	var result []int
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i+1 {
// 			result = append(result, arr[i], i+1)
// 		}
// 	}
// 	return result
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	var result []int
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i+1 {
// 			result = append(result, arr[i], i+1) // arr[i] is for duplicate and i+1 is for missing
// 			break
// 		}
// 	}
// 	return result
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }
