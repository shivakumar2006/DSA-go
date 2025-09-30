// Peak index in mountain array
// https://leetcode.com/problems/peak-index-in-a-mountain-array/description/

package main

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {
		mid := start + (end-start)/2
		if arr[mid] > arr[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func main() {
	arr := []int{0, 1, 0}
	fmt.Println("peak index is : ", peakIndexInMountainArray(arr))
}
