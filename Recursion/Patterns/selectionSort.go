package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 1}
	selection(arr, len(arr), 0, 0)
	fmt.Println(arr)
}

func selection(arr []int, r int, c int, max int) {
	if r == 0 {
		return
	}

	if c < r {
		if arr[c] > arr[max] {
			selection(arr, r, c+1, c)
		} else {
			selection(arr, r, c+1, max)
		}
	} else {
		temp := arr[max]
		arr[max] = arr[r-1]
		arr[r-1] = temp
		selection(arr, r-1, 0, 0)
	}
}

// func selection(arr []int, r, c int, max int) {
// 	if r == 0 {
// 		return
// 	}

// 	if c < r {
// 		if arr[c] > arr[max] {
// 			selection(arr, r, c+1, c)
// 		} else {
// 			selection(arr, r, c+1, max)
// 		}
// 	} else {
// 		temp := arr[max]
// 		arr[max] = arr[r-1]
// 		arr[r-1] = temp
// 		selection(arr, r-1, 0, 0)
// 	}
// }

// func selection(arr []int, r, c int, max int) {
// 	if r == 0 {
// 		return
// 	}

// 	if c < r {
// 		if arr[c] > arr[max] {
// 			selection(arr, r, c+1, c)
// 		} else {
// 			selection(arr, r, c+1, max)
// 		}
// 	} else {
// 		temp := arr[max]
// 		arr[max] = arr[r-1]
// 		arr[r-1] = temp
// 		selection(arr, r-1, 0, 0)
// 	}
// }
