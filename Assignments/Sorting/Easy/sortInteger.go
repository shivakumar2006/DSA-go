// sort integers by the number of 1 Bits
// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/description/

package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		countI := bits.OnesCount(uint(arr[i]))
		countJ := bits.OnesCount(uint(arr[j]))

		if countI == countJ {
			return arr[i] < arr[j]
		}
		return countI < countJ
	})
	return arr
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sortByBits(arr))
}
