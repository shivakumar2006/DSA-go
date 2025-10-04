// Find in Mountain Array
// https://leetcode.com/problems/find-in-mountain-array/description/

package main

import "fmt"

func findInMountainArray(arr []int, tar int) int {
	peak := findPeak(arr)

	index := binarySearch(arr, tar, 0, peak, true)
	if index != -1 {
		return index
	}
	return binarySearch(arr, tar, peak+1, len(arr)-1, false)
}

func findPeak(arr []int) int {
	start, end := 0, len(arr)-1
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

func binarySearch(arr []int, tar int, start, end int, asc bool) int {
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			return mid
		}

		if asc {
			if arr[mid] < tar {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if arr[mid] < tar {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 3, 1}
	tar := 3
	ans := findInMountainArray(arr, tar)
	fmt.Println(ans)
}
