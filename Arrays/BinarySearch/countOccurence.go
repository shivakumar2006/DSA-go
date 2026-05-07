package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 2, 3}
	tar := 2
	res := count(arr, tar)
	fmt.Printf("The count is : %d", res)
}

func count(arr []int, tar int) int {
	first := findFirst(arr, tar)
	if first == -1 {
		return 0
	}
	last := findLast(arr, tar)
	return last - first + 1
}

func findFirst(arr []int, tar int) int {
	start, end := 0, len(arr)-1
	result := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			result = mid
			end = mid - 1
		} else if arr[mid] < tar {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result
}

func findLast(arr []int, tar int) int {
	start, end := 0, len(arr)-1
	result := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			result = mid
			start = mid + 1
		} else if arr[mid] < tar {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result
}

// func count(arr []int, tar int) int {
// 	first := findFirst(arr, tar)
// 	if first == -1 {
// 		return 0
// 	}
// 	last := findLast(arr, tar)
// 	return last - first + 1
// }

// func findFirst(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	result := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if arr[mid] == tar {
// 			result = mid
// 			end = mid - 1
// 		} else if arr[mid] < tar {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return result
// }

// func findLast(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	result := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if arr[mid] == tar {
// 			result = mid
// 			start = mid + 1
// 		} else if arr[mid] < tar {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return result
// }
