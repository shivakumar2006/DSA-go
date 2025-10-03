// Find the duplicate number
// https://leetcode.com/problems/find-the-duplicate-number/

package main

import "fmt"

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	low := 1
	high := n

	for low < high {
		mid := low + (high-low)/2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}

// func findDuplicate(nums []int) int {
// 	n := len(nums) - 1
// 	low := 1
// 	high := n

// 	for low < high {
// 		mid := low + (high-low)/2
// 		count := 0
// 		for _, num := range nums {
// 			if num <= mid {
// 				count++
// 			}
// 		}

// 		if count > mid {
// 			high = mid
// 		} else {
// 			low = mid + 1
// 		}
// 	}
// 	return low
// }
