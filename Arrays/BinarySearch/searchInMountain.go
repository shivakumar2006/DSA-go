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
		fmt.Println("Target not found")
	}
}

func findPeak(arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		mid := start + (end-start)/2
		if arr[mid] > arr[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func searchInMountain(arr []int, tar int) int {
	peak := findPeak(arr)

	// search in ascending part
	index := binarySearch(arr, tar, 0, peak, true)
	if index != -1 {
		return index
	}

	// searcn in descending oreder
	return binarySearch(arr, tar, peak+1, len(arr)-1, false)
}

// order-agnostic binarySarch
func binarySearch(arr []int, tar int, start int, end int, asc bool) int {
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			return mid
		}
		if asc {
			if tar < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if tar > arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{1, 2, 3, 5, 6, 4, 2, 1}
// 	tar := 5
// 	res := searchInMountain(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The index of the target is : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func findPeak(arr []int) int {
// 	start, end := 0, len(arr)-1
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

// func searchInMountain(arr []int, tar int) int {
// 	peak := findPeak(arr)

// 	// search in ascending part
// 	index := binarySearch(arr, tar, 0, peak, true)
// 	if index != -1 {
// 		return index
// 	}

// 	// searcn in descending oreder
// 	return binarySearch(arr, tar, peak+1, len(arr)-1, false)
// }

// // order-agnostic binarySarch
// func binarySearch(arr []int, tar int, start int, end int, asc bool) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if arr[mid] == tar {
// 			return mid
// 		}
// 		if asc {
// 			if tar < arr[mid] {
// 				end = mid - 1
// 			} else {
// 				start = mid + 1
// 			}
// 		} else {
// 			if tar > arr[mid] {
// 				end = mid - 1
// 			} else {
// 				start = mid + 1
// 			}
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{1, 2, 3, 5, 6, 4, 2, 1}
// 	tar := 5
// 	res := searchInMountain(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The index of the target is : %d\n", res)
// 	} else {
// 		fmt.Println("target npt found")
// 	}
// }

// func searchInMountain(arr []int, tar int) int {
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
// 	return binarySearch(arr, tar)
// }

// func binarySearch(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	for l < r {
// 		mid := l + (r-l)/2
// 		if x == arr[mid] {
// 			return mid
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return -1
// }
