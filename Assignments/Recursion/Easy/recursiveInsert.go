// REcursive Insertion sort
// https://www.geeksforgeeks.org/dsa/recursive-insertion-sort/

package main

import "fmt"

func insertionRecursiveSort(arr []int, n int) {
	if n <= 1 {
		return
	}

	insertionRecursiveSort(arr, n-1)

	last := arr[n-1]
	j := n - 2

	for j >= 0 && arr[j] > last {
		arr[j+1] = arr[j]
		j--
	}

	arr[j+1] = last
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	insertionRecursiveSort(arr, len(arr))
	fmt.Println(arr)
}
