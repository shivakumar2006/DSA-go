// First missing positive integer
// https://leetcode.com/problems/first-missing-positive/description/

package main

import "fmt"

func firstMissingPositive(arr []int) int {
	// place every num on it's correct position
	for i := 0; i < len(arr); i++ {
		for arr[i] > 0 && arr[i] <= len(arr) && arr[i] != arr[arr[i]-1] {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		}
	}

	// find the first index where nums[i] != i + 1
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			return i + 1
		}
	}

	return len(arr) + 1
}

func main() {
	arr := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(arr))
}
