// single element in sorted array

package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	res := search(arr)
	fmt.Println(res)
}

func search(arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		mid := start + (end-start)/2
		if mid%2 != 0 {
			mid--
		}

		if arr[mid] == arr[mid+1] {
			start = mid + 2
		} else {
			end = mid
		}
	}
	return arr[start]
}

// func search(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start < end {
// 		mid := start + (end-start)/2
// 		if mid%2 == 1 {
// 			mid--
// 		}

// 		if arr[mid] == arr[mid+1] {
// 			start = mid + 2
// 		} else {
// 			end = mid
// 		}
// 	}
// 	return arr[start]
// }
