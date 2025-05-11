// Find the peak index in mountain array

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 4, 5, 3, 2, 1}
	res := peak(arr)
	fmt.Printf("Peak index is : %d\n", res)
}

func peak(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {
		mid := start + (end-start)/2 // (start + end) / 2
		if arr[mid] > arr[mid+1] {
			// means we are in descreasing array
			end = mid
		} else {
			// means we are in increasing array
			start = mid + 1
		}
	}
	return start
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{0, 1, 0}
// 	res := mountainArray(arr)
// 	fmt.Printf("Peak index is : %d\n", res)
// }

// func mountainArray(arr []int) int {
// 	start := 0
// 	end := len(arr) - 1

// 	for start < end {
// 		mid := start + (end-start)/2
// 		if arr[mid] > arr[mid+1] {
// 			end = mid
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return start
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{0, 2, 3, 5, 3, 2, 1}
// 	peak := mountainArray(arr)
// 	fmt.Printf("Peak index is : %d\n", peak)
// }

// func mountainArray(arr []int) int {
// 	start := 0
// 	end := len(arr) - 1

// 	for start < end {
// 		mid := start + (end-start)/2
// 		if arr[mid] > arr[mid+1] {
// 			end = mid
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return start
// }
