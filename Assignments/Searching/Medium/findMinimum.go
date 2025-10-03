// Find minimum in rotated sorted array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

package main

import "fmt"

func main() {
	arr := []int{3, 4, 5, 1, 2}
	fmt.Println("Minimm element : ", findMin(arr))
}

func findMin(arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		mid := start + (end-start)/2
		if arr[mid] > arr[end] {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return arr[start]
}

// func findMin(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start < end {
// 		mid := start + (end-start)/2
// 		if arr[mid] > arr[end] {
// 			start = mid + 1
// 		} else {
// 			end = mid
// 		}
// 	}
// 	return arr[start]
// }
