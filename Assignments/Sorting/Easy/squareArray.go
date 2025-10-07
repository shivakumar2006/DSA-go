// Square of a sorted arrays
// https://leetcode.com/problems/squares-of-a-sorted-array/description/

package main

import "fmt"

func sortedSquares(nums []int) []int {
	set := []int{}
	for _, val := range nums {
		set = append(set, val*val)
	}
	for i := 1; i < len(set); i++ {
		j := i
		for j > 0 && set[j] < set[j-1] {
			swap(set, j, j-1)
			j--
		}
	}

	return set
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
