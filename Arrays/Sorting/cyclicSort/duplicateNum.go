// Find the duplicate numbers

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 4}
	res := sort(arr)
	fmt.Println(res)
}

func sort(arr []int) int {
	i := 0
	for i < len(arr) {
		correct := arr[i] - 1
		if arr[i] != arr[correct] {
			swap(arr, i, correct)
		} else {
			if i != correct {
				return arr[i]
			}
			i++
		}
	}
	return -1
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// func sort(arr []int) int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			if i != correct {
// 				return arr[i] // duplicate found
// 			}
// 			i++
// 		}
// 	}
// 	return -1
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }
