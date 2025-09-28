// Search Insert Position
// https://leetcode.com/problems/search-insert-position/description/

package main

import "fmt"

func searchInsert(nums []int, tar int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == tar {
			return mid
		} else if nums[mid] < tar {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func main() {
	nums := []int{1, 3, 5, 6}
	tar := 3
	fmt.Println("insert position is : ", searchInsert(nums, tar))
}
