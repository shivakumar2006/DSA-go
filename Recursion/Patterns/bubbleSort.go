package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 1}
	bubble(arr, len(arr)-1, 0)
	fmt.Println(arr)
}

func bubble(arr []int, r, c int) {
	if r == 0 {
		return
	}

	if c < r {
		if arr[c] > arr[c+1] {
			temp := arr[c]
			arr[c] = arr[c+1]
			arr[c+1] = temp
		}
		bubble(arr, r, c+1)
	} else {
		bubble(arr, r-1, 0)
	}
}

// func bubble(arr []int, r, c int) {
// 	if r == 0 {
// 		return
// 	}

// 	if c < r {
// 		if arr[c] > arr[c+1] {
// 			temp := arr[c]
// 			arr[c] = arr[c+1]
// 			arr[c+1] = temp
// 		}
// 		bubble(arr, r, c+1)
// 	} else {
// 		bubble(arr, r-1, 0)
// 	}
// }

// func bubble(arr []int, r int, c int) {
// 	if r == 0 {
// 		return
// 	}

// 	if c < r {
// 		if arr[c] > arr[c+1] {
// 			temp := arr[c]
// 			arr[c] = arr[c+1]
// 			arr[c+1] = temp
// 		}
// 		bubble(arr, r, c+1)
// 	} else {
// 		bubble(arr, r-1, 0)
// 	}
// }
