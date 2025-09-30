// Rotated sorted array
// https://leetcode.com/problems/search-in-rotated-sorted-array/description/

package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	ans := search(arr, target)
	if ans != 0 {
		fmt.Println("the target element is found at the index : ", ans)
	} else {
		fmt.Println("not found")
	}
}

func search(arr []int, tar int) int {
	pivot := findPivot(arr)

	if pivot == -1 {
		binarySearch(arr, tar, 0, len(arr)-1)
	}

	if tar == arr[pivot] {
		return pivot
	}

	if tar >= arr[0] {
		binarySearch(arr, tar, 0, pivot)
	}
	return binarySearch(arr, tar, pivot+1, len(arr)-1)
}

func findPivot(arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		mid := start + (end-start)/2
		if mid < end && arr[mid] > arr[mid+1] {
			return mid
		} else if mid > start && arr[mid] < arr[mid-1] {
			return mid - 1
		} else if arr[mid] >= arr[start] {
			start = mid + 1
		} else {
			end = end - 1
		}
	}
	return -1
}

func binarySearch(arr []int, tar int, start, end int) int {
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			return mid
		} else if arr[mid] < tar {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
