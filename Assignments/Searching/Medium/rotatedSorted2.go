// rotated sorted 2
// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/

// Rotated sorted array
// https://leetcode.com/problems/search-in-rotated-sorted-array/description/

package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 0, 1, 2}
	target := 0
	ans := search(arr, target)
	fmt.Println("the target element is found at the index : ", ans)
}

func search(arr []int, tar int) bool {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			return true
		}

		if arr[start] == arr[mid] && arr[start] == arr[end] {
			start++
			end--
		} else if arr[start] <= arr[mid] { // left sorted
			if arr[start] <= tar && tar < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else { // right sorted
			if arr[mid] < tar && tar <= arr[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}

// func search(arr []int, tar int) bool {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if arr[mid] == tar {
// 			return true
// 		}

// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			start++
// 			end--
// 		} else if arr[start] <= arr[mid] { // means left side already sorted
// 			if arr[start] <= tar && tar < arr[mid] {
// 				end = mid - 1
// 			} else {
// 				start = mid + 1
// 			}
// 		} else {
// 			if arr[mid] < tar && tar <= arr[end] { // means right side already soerted
// 				start = mid + 1
// 			} else {
// 				end = mid - 1
// 			}
// 		}
// 	}
// 	return false
// }
