// Find the mountain array...

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 5, 6, 4, 2, 1}
	tar := 5
	res := searchInMountain(arr, tar)
	if res != -1 {
		fmt.Printf("The index of the target is : %d\n", res)
	} else {
		fmt.Println("target npt found")
	}
}

func searchInMountain(arr []int, tar int) int {
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
	return binarySearch(arr, tar)
}

func binarySearch(arr []int, x int) int {
	l, r := 0, len(arr)-1
	for l < r {
		mid := l + (r-l)/2
		if x == arr[mid] {
			return mid
		} else if x < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
