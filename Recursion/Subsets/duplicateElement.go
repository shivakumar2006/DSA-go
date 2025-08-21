package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 2}
	fmt.Println(subsets(arr))
}

func subsets(arr []int) [][]int {
	sort.Ints(arr)
	result := [][]int{{}}
	start, end := 0, 0

	for i := 0; i < len(arr); i++ {
		start = 0
		if i > 0 && arr[i] == arr[i-1] {
			start = end + 1
		}
		end = len(result) - 1
		n := len(result)
		for j := start; j < n; j++ {
			newSubset := append([]int{}, result[j]...)
			newSubset = append(newSubset, arr[i])
			result = append(result, newSubset)
		}
	}
	return result
}

// func subsets(arr []int) [][]int {
// 	sort.Ints(arr)
// 	result := [][]int{{}}
// 	start, end := 0, 0

// 	for i := 0; i < len(arr); i++ {
// 		start = 0
// 		if i > 0 && arr[i] == arr[i-1] {
// 			start = end + 1
// 		}
// 		end = len(result) - 1
// 		n := len(result)
// 		for j := start; j < n; j++ {
// 			newSubset := append([]int{}, result[j]...)
// 			newSubset = append(newSubset, arr[i])
// 			result = append(result, newSubset)
// 		}
// 	}
// 	return result
// }

// func subsets(arr []int) [][]int {
// 	sort.Ints(arr)
// 	result := [][]int{{}}
// 	start, end := 0, 0

// 	for i := 0; i < len(arr); i++ {
// 		start = 0
// 		if i > 0 && arr[i] == arr[i-1] {
// 			start = end + 1
// 		}

// 		end = len(result) - 1
// 		n := len(result)
// 		for j := start; j < n; j++ {
// 			newSubset := append([]int{}, result[j]...)
// 			newSubset = append(newSubset, arr[i])
// 			result = append(result, newSubset)
// 		}
// 	}
// 	return result
// }

// func subsets(arr []int) [][]int {
// 	sort.Ints(arr) // handle duplicates
// 	result := [][]int{{}}
// 	start, end := 0, 0

// 	for i := 0; i < len(arr); i++ {
// 		start = 0
// 		// if current element is same as previous, only then starts from end+1
// 		if i > 0 && arr[i] == arr[i-1] {
// 			start = end + 1
// 		}
// 		end = len(result) - 1
// 		n := len(result)
// 		for j := start; j < n; j++ {
// 			// coppy existing subsets
// 			newSubset := append([]int{}, result[j]...)
// 			newSubset = append(newSubset, arr[i])
// 			result = append(result, newSubset)
// 		}
// 	}
// 	return result
// }
