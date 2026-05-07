// Find minimum in rotated sorted array

package main

import "fmt"

func main() {
	arr := []int{3, 4, 5, 1, 2}
	fmt.Printf("the minimim element is found at the index : %d\n", search(arr))
}

func search(arr []int) int {
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

// func search(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start < end {
// 		mid := start + (end-start)/2
// 		if arr[mid] > arr[end] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return arr[start]
// }

// func search(arr []int) int {
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
