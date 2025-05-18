// Split array largest sum ...

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{7, 2, 5, 10, 8}
	m := 2
	fmt.Println(split(nums, m)) // Output should be 18
}

func split(nums []int, m int) int {
	start, end := 0, 0
	for i := 0; i < len(nums); i++ {
		start = int(math.Max(float64(start), float64(nums[i])))
		end += nums[i]
	}

	for start < end {
		mid := start + (end-start)/2

		sum := 0
		pieces := 1

		for i := 0; i < len(nums); i++ {
			if sum+nums[i] > mid {
				sum = nums[i]
				pieces++
			} else {
				sum += nums[i]
			}
		}

		if pieces > m {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}
