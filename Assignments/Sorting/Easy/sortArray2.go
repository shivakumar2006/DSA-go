// Sort Array By Parity II
// https://leetcode.com/problems/sort-array-by-parity-ii/description/

package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	n := len(nums)

	for even < n && odd < n {
		if nums[even]%2 == 0 {
			even += 2
		} else if nums[odd]%2 == 1 {
			odd += 2
		} else {
			nums[even], nums[odd] = nums[odd], nums[even]
		}
	}
	return nums
}

func main() {
	nums := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(nums))
}
