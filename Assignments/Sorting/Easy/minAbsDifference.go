// Minimum Absolute Difference
// https://leetcode.com/problems/minimum-absolute-difference/description/

package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	minDiff := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	result := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minDiff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}

	return result
}

func main() {
	arr := []int{4, 2, 1, 3}
	fmt.Println(minimumAbsDifference(arr))
}

// func minimumAbsDifference(arr []int) [][]int {
// 	sort.Ints(arr)

// 	// find the minimum absolute difference
// 	minDiff := arr[1] - arr[0]
// 	for i := 1; i < len(arr)-1; i++ {
// 		diff := arr[i+1] - arr[i]
// 		if diff < minDiff {
// 			minDiff = diff
// 		}
// 	}

// 	// collect all pairs with that min difference
// 	result := [][]int{}
// 	for i := 0; i < len(arr)-1; i++ {
// 		if arr[i+1]-arr[i] == minDiff {
// 			result = append(result, []int{arr[i], arr[i+1]})
// 		}
// 	}

// 	return result
// }
