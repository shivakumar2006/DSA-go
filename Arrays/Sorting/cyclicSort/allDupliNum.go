package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 4, 5, 6, 6}
	res := sort(arr)
	fmt.Println(res)
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
// 	var duplicate []int
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != i+1 {
// 			if !contains(duplicate, arr[i]) {
// 				duplicate = append(duplicate, arr[i])
// 			}
// 		}
// 	}
// 	return duplicate
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func contains(arr []int, value int) bool {
// 	for _, v := range arr {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }
