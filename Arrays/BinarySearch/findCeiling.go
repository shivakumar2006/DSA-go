// Find the ceiling of the number of the middle or smallest number greater >= the middle in sorted array

package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 5, 6, 9, 14}
	tar := 12
	res := findCeiling(arr, tar)
	if res != -1 {
		fmt.Printf("The ceiling of the target number is found at the index : %d\n", res)
	} else {
		fmt.Println("Target is not found")
	}
}

func findCeiling(arr []int, x int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if x == arr[mid] {
			return mid
		} else if x < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if start < len(arr) {
		return start
	}
	return -1
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 14}
// 	tar := 8
// 	res := findCeiling(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The ceiling of the target number is found at the index : %d\n", res)
// 	} else {
// 		fmt.Println("target number is not found")
// 	}
// }

// func findCeiling(arr []int, x int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if x == arr[mid] {
// 			return mid
// 		} else if x < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	if start < len(arr) {
// 		return start
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 14}
// 	tar := 7
// 	res := findCeiling(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The ceiling of the target number is found at the index of : %d\n", res)
// 	} else {
// 		fmt.Println("The ceiling of the tsrget number is not found...")
// 	}
// }

// func findCeiling(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			return mid
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	if l < len(arr) {
// 		return l
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 14, 16, 18}
// 	tar := 15
// 	res := binarySearch(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The ceiling of the target number is found at the index of : %d\n", res)
// 	} else {
// 		fmt.Println("The ceiling of the target number is not found...")
// 	}
// }

// func binarySearch(arr []int, x int) int {
// 	l := 0
// 	r := len(arr) - 1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			return mid
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	if l < len(arr) {
// 		return l
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 14, 16, 18}
// 	tar := 15
// 	res := findCeiling(arr, tar)
// 	fmt.Printf("The ceiling of the target number if found at the index of : %d\n", res)
// }

// func findCeiling(arr []int, x int) int {
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
// 	if l < len(arr) {
// 		return l
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 14, 16, 18}
// 	tar := 15
// 	res := findCeiling(arr, tar)
// 	fmt.Printf("The Ceiling of the target number is found at the index of : %d\n", res)
// }

// func findCeiling(arr []int, x int) int {
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
// 	if l < len(arr) {
// 		return l
// 	}
// 	return -1
// }
