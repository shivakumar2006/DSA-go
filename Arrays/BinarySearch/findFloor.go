package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 5, 6, 9, 11, 14, 16, 20}
	tar := 15
	res := findFloor(arr, tar)
	fmt.Printf("The floor of the taret number is found at the index of : %d\n", res)
}

func findFloor(arr []int, x int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == x {
			return mid
		} else if x < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if r >= 0 {
		return r
	}
	return -1
}
