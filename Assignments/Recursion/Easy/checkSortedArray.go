// check if array is sorted or not
// https://www.geeksforgeeks.org/dsa/program-check-array-sorted-not-iterative-recursive/

package main

import "fmt"

func isSorted(arr []int, n int) bool {
	if n == 1 || n == 0 {
		return true
	}

	return arr[n-1] >= arr[n-2] && isSorted(arr, n-1)
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	n := len(arr)
	if isSorted(arr, n) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
