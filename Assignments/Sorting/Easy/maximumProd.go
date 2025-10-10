// maximum product of two element in the array
// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/description/

package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max1, max2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max1 {
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 {
			max2 = nums[i]
		}
	}

	calc := (max1 - 1) * (max2 - 1)
	return calc
}

func main() {
	nums := []int{3, 4, 5, 2}
	fmt.Println(maxProduct(nums))
}
