// Largest Perimeter Traingle
// https://leetcode.com/problems/largest-perimeter-triangle/description/

package main

import "fmt"

func largestPerimeter(nums []int) int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] > nums[j-1]; j-- {
			swap(nums, j, j-1)
		}
	}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1]+nums[i+2] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	nums := []int{2, 1, 2}
	fmt.Println(largestPerimeter(nums))
}
