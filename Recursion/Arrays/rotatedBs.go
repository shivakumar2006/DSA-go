// searching in rotated binarySearch using recursion

package main

import "fmt"

func main() {
	arr := []int{5, 6, 7, 8, 9, 1, 2, 3}
	fmt.Println(search(arr, 8, 0, len(arr)-1))
}

func search(arr []int, tar int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if arr[mid] == tar {
		return mid
	}

	if arr[start] <= arr[mid] {
		if tar >= arr[start] && tar <= arr[mid] {
			return search(arr, tar, end, mid-1)
		} else {
			return search(arr, tar, start, mid+1)
		}
	}

	if arr[start] >= arr[mid] && tar <= arr[end] {
		return search(arr, tar, start, mid+1)
	}
	return search(arr, tar, end, mid-1)
}
