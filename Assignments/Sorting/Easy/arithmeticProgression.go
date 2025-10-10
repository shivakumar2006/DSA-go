// Can make arithmetic progression from sequences
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/description/

package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	// for i := 0; i < len(arr)-1; i++ {
	// 	for j := 0; j < len(arr)-i-1; j++ {
	// 		if arr[j] < arr[j+1] {
	// 			swap(arr, j, j+1)
	// 		}
	// 	}
	// }

	sort.Ints(arr)

	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

// func swap(arr []int, i, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

func main() {
	arr := []int{3, 5, 1}
	fmt.Println(canMakeArithmeticProgression(arr))
}
