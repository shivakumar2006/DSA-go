// bs using recursion

package main

import "fmt"

func bs(arr []int, tar int, start, end int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if arr[mid] == tar {
		return mid
	}

	if arr[mid] < tar {
		return bs(arr, tar, mid+1, end)
	}
	return bs(arr, tar, start, mid-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	tar := 3
	fmt.Println(bs(arr, tar, 0, len(arr)-1))
}
