// Recursive Bubble Sort
// https://www.geeksforgeeks.org/dsa/recursive-bubble-sort/

package main

import "fmt"

func bubbleSort(arr []int, n int) {
	if n == 1 {
		return
	}

	count := 0
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
			count++
		}
	}

	if count == 0 {
		return
	}

	bubbleSort(arr, n-1)
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr, len(arr))
	fmt.Print("Sorted Array : ")
	fmt.Println(arr)
}
