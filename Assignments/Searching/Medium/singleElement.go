// Single Element in a Sorted aray
// https://leetcode.com/problems/single-element-in-a-sorted-array/description/

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	fmt.Println("singleNonDuplicate : ", singleNonDuplicate(arr))
}

func singleNonDuplicate(arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		mid := start + (end-start)/2
		if mid%2 == 1 {
			mid = mid - 1
		}
		if arr[mid] == arr[mid+1] {
			start = mid + 2
		} else {
			end = mid
		}
	}
	return arr[start]
}
