// Find all numbers dissappeared in the array..., here is again from 1 to n then use cyclic Sort
// Google question

package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 3, 2, 0, 6}
	res := sort(arr)
	fmt.Println(res)
}

func sort(arr []int) []int {
	i := 0
	for i < len(arr) {
		correct := arr[i] - 1
		if arr[i] > 0 && correct < len(arr) && arr[i] != arr[correct] {
			swap(arr, i, correct)
		} else {
			i++
		}
	}
	var missing []int
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			missing = append(missing, i+1)
		}
	}
	return missing
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] > 0 && correct < len(arr) && arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	var missing []int
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i+1 {
// 			missing = append(missing, i+1)
// 		}
// 	}
// 	return missing
// }

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] > 0 && correct < len(arr) && arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	var missing []int
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i+1 {
// 			missing = append(missing, i+1)
// 		}
// 	}
// 	return missing
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] > 0 && correct < len(arr) && arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}

// 	var missing []int
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i+1 {
// 			missing = append(missing, i+1)
// 		}
// 	}
// 	return missing
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] > 0 && correct < len(arr) && arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}

// 	var missing []int
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i+1 {
// 			missing = append(missing, i+1)
// 		}
// 	}
// 	return missing
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }
