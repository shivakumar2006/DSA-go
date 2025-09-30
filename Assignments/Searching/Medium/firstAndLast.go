// find first and last
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/

package main

import "fmt"

func main() {
	arr := []int{5, 7, 7, 8, 8, 10}
	tar := 8
	ans := searchRange(arr, tar)
	if ans != nil {
		fmt.Printf("the first and last index is : [%d, %d] ", ans[0], ans[1])
	} else {
		fmt.Println("not found")
	}
}

func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)

	if first == -1 && last == -1 {
		return []int{-1, -1}
	}

	return []int{first, last}
}

func findFirst(arr []int, tar int) int {
	start, end := 0, len(arr)-1
	res := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			res = mid
			end = mid - 1
		} else if tar < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return res
}

func findLast(arr []int, tar int) int {
	start, end := 0, len(arr)-1
	res := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			res = mid
			start = mid + 1
		} else if tar < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return res
}
