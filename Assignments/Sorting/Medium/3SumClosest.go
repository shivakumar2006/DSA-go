// 3 sum closest
// https://leetcode.com/problems/3sum-closest/description/

package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, tar int) int {
	sort.Ints(nums)
	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if math.Abs(float64(sum-tar)) < math.Abs(float64(closestSum-tar)) {
				closestSum = sum
			}

			if sum < tar {
				left++
			} else if sum > tar {
				right--
			} else {
				return sum
			}
		}
	}
	return closestSum
}

func main() {
	nums := []int{-1, 2, 1, -4}
	tar := 1
	fmt.Println(threeSumClosest(nums, tar))
}
