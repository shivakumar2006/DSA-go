// Frequency of the most frequent element
// https://leetcode.com/problems/frequency-of-the-most-frequent-element/description/

package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	left, total, res := 0, 0, 1

	for right := 0; right < len(nums); right++ {
		total += nums[right]

		// while cost > k, shrink window
		for (right-left+1)*nums[right]-total > k {
			total -= nums[left]
			left++
		}

		// update max frequency
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 4}
	k := 5
	fmt.Println(maxFrequency(nums, k))
}
