// Find the missing number in the array and the range of the array is 0 to n

// NOTE := when we see the range is 0 to N and 1 to N, then we use cyclic sort...

// Amazon question

package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 0, 1}
	res := sort(arr)
	fmt.Println(res)
}

func sort(arr []int) int {
	i := 0
	for i < len(arr) {
		correct := arr[i]
		if correct < len(arr) && arr[i] != arr[correct] {
			swap(arr, i, correct)
		} else {
			i++
		}
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			return i
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
// 		correct := arr[i]
// 		if correct < len(arr) && arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i {
// 			return i
// 		}
// 	}
// 	return len(arr)
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }
