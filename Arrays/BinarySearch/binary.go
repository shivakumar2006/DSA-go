package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 5, 3, 6, 8, 10, 20}
	fmt.Println("the index number if binary search is : ", binarySearch(arr, 6))
}

func binarySearch(arr []int, x int) int {
	l := 0
	r := len(arr)
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
	return -1
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 5, 3, 6, 8, 10, 20}
// 	fmt.Println(binarySearch(arr, 20))
// }

// func binarySearch(arr []int, x int) int {
// 	l := 0
// 	r := len(arr) - 1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if arr[mid] == x {
// 			return mid
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 5, 6, 3, 69, 78, 42}
// 	fmt.Println(binarySearch(arr, 69))
// }

// func binarySearch(arr []int, x int) int {
// 	l := 0
// 	r := len(arr) - 1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if arr[mid] == x {
// 			return mid
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return 1
// }
