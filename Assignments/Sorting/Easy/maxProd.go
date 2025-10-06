// Maximum Product of 3 numbers
// https://leetcode.com/problems/maximum-product-of-three-numbers/

package main

import (
	"fmt"
	"sort"
)

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	prod1 := nums[n-1] * nums[n-2] * nums[n-3]
	prod2 := nums[0] * nums[1] * nums[n-1]

	if prod1 > prod2 {
		return prod1
	}
	return prod2
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(maximumProduct(nums))
}
