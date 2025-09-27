// Two Sum 2 - input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

package main

import "fmt"

func main() {
	num := []int{2, 7, 11, 15}
	tar := 9
	fmt.Println("The index of the target sum is : ", twoSum(num, tar))
}

func twoSum(num []int, tar int) []int {
	left, right := 0, len(num)-1
	sum := 0
	for left < right {
		sum = num[left] + num[right]
		if sum < tar {
			left++
		} else if sum > tar {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

// func twoSum(num []int, tar int) []int {
// 	left := 0
// 	right := len(num) - 1
// 	sum := 0
// 	for left < right {
// 		sum = num[left] + num[right]
// 		if sum < tar {
// 			left++
// 		} else if sum > tar {
// 			right--
// 		} else {
// 			return []int{left + 1, right + 1}
// 		}
// 	}
// 	return []int{}
// }
