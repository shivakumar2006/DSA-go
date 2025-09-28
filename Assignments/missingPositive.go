// Kth Missing Positive number
// https://leetcode.com/problems/kth-missing-positive-number/description/

package main

import "fmt"

func main() {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	fmt.Println("missing positive integer is : ", findKthPositive(arr, k))
}

func findKthPositive(arr []int, k int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2
		missing := arr[mid] - (mid + 1)
		if missing < k {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start + k
}
