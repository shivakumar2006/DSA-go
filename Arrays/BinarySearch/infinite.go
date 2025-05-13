// find the target element in a infinite sorted array...
// this is the sorted array so, obviously we use binary search...
// this is amazon interview question...

package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 5, 6, 8, 9, 10, 13, 16, 20, 30, 35, 40}
	tar := 20
	res := search(arr, tar)
	if res != -1 {
		fmt.Printf("Target %d found at the index : %d\n", tar, res)
	} else {
		fmt.Println("Target not found")
	}
}

func search(arr []int, tar int) int {
	start, end := 0, len(arr)-1
	for tar > arr[end] {
		newStart := end + 1
		end = end + (end-start+1)*2
		if end >= len(arr) {
			end = len(arr) - 1
		}
		start = newStart
	}
	return binarySearch(arr, tar, start, end)
}
func binarySearch(arr []int, tar int, start int, end int) int {
	for start < end {
		mid := start + (end-start)/2
		if tar == arr[mid] {
			return mid
		} else if tar < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 5, 6, 8, 9, 10, 13, 16, 20, 30, 35, 40}
// 	tar := 20
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("Target  %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	start, end := 0, 1

// 	for tar > arr[end] {
// 		newStart := end + 1
// 		end = end + (end-start+1)*2
// 		if end >= len(arr) {
// 			end = len(arr) - 1
// 		}
// 		start = newStart
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 5, 6, 8, 9, 10, 13, 16, 20, 30, 35, 40}
// 	tar := 16
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("Tatget %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	start, end := 0, 1

// 	for tar > arr[end] {
// 		newStart := end + 1
// 		end = end + (end-start+1)*2
// 		if end >= len(arr) {
// 			end = len(arr) - 1
// 		}
// 		start = newStart
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start < end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 5, 6, 8, 9, 10, 13, 16, 20, 30, 35, 40}
// 	tar := 16
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("Target %d found ast the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	start := 0
// 	end := 1

// 	for tar > arr[end] {
// 		newStart := end + 1
// 		end = end + (end-start+1)*2
// 		if end >= len(arr) {
// 			end = len(arr) - 1
// 		}
// 		start = newStart
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 8, 9, 11, 13, 16, 20, 30, 35, 40}
// 	tar := 35
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("Target %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	start := 0
// 	end := 1

// 	for tar > arr[end] {
// 		temp := end + 1
// 		end = end + (end-start+1)*2

// 		if end >= len(arr) {
// 			end = len(arr) - 1
// 		}

// 		start = temp
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 8, 9, 11, 13, 16, 20, 30, 35, 40}
// 	tar := 35
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("Target %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	start := 0
// 	end := 1

// 	for tar > arr[end] {
// 		temp := end + 1
// 		end = end + (end-start+1)*2

// 		if end >= len(arr) { // for prevent out of bound error
// 			end = len(arr) - 1
// 		}

// 		start = temp
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 8, 9, 11, 13, 16, 20, 30, 35, 40}
// 	tar := 30
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("Target %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	start := 0
// 	end := 1

// 	for tar > arr[end] {
// 		temp := end + 1
// 		end = end + (end-start+1)*2

// 		if end >= len(arr) {
// 			end = len(arr) - 1
// 		}

// 		start = temp
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// var arr = []int{2, 4, 5, 6, 8, 9, 11, 13, 16, 20, 30, 35, 40}

// func main() {
// 	tar := 30
// 	res := search(tar)
// 	if res != -1 {
// 		fmt.Printf("Target %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(tar int) int {
// 	start := 0
// 	end := 1

// 	for tar > arr[end] {
// 		temp := end + 1
// 		end = end + (end-start+1)*2

// 		if end >= len(arr) {
// 			end = len(arr) - 1
// 		}

// 		start = temp
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// var arr = []int{2, 4, 5, 6, 8, 9, 11, 13, 16, 20, 30, 35, 40}

// func main() {
// 	tar := 30
// 	res := search(tar)
// 	if res != -1 {
// 		fmt.Printf("Target %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found...")
// 	}
// }

// func search(tar int) int {
// 	start := 0
// 	end := 1

// 	for tar > arr[end] {
// 		temp := end + 1
// 		end = end + (end-start+1)*2

// 		if end >= len(arr) { // prevent out of bounds error
// 			end = len(arr) - 1
// 		}

// 		start = temp
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// var arr = []int{2, 4, 5, 7, 9, 11, 13, 16, 20, 25, 30, 35, 40}

// func main() {
// 	tar := 25
// 	res := search(tar)
// 	if res != -1 {
// 		fmt.Printf("Target %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found...")
// 	}
// }

// func search(tar int) int {
// 	start := 0
// 	end := 1

// 	for tar > arr[end] {
// 		temp := end + 1
// 		end = end + (end-start+1)*2
// 		start = temp

// 		// Add boundary check to avoid panic
// 		if end >= len(arr) {
// 			end = len(arr) - 1
// 			break
// 		}
// 	}
// 	return binarySearch(arr, tar, start, end)
// }

// func binarySearch(arr []int, tar int, start int, end int) int {
// 	for start <= end {
// 		mid := start + (end-start)/2

// 		if tar < arr[mid] {
// 			end = mid - 1
// 		} else if tar > arr[mid] {
// 			start = mid + 1
// 		} else {
// 			return mid
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// var arr = []int{2, 4, 5, 7, 9, 11, 13, 16, 20, 25, 30, 35, 40}

// func main() {
// 	tar := 25
// 	res := search(tar)
// 	if res != -1 {
// 		fmt.Printf("Target %d found at the index : %d\n", tar, res)
// 	} else {
// 		fmt.Println("Target not found...")
// 	}
// }

// func search(tar int) int {
// 	l := 0
// 	r := 1

// 	for get(r) < tar {
// 		l = r
// 		r *= 2
// 	}

// 	for l <= r {
// 		mid := l + (r-l)/2
// 		val := get(mid)
// 		if val == tar {
// 			return mid
// 		} else if val < tar {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return -1
// }

// func get(res int) int {
// 	if res >= len(arr) {
// 		return int(^uint(0) >> 1) or return math.MaxInt
// 	}
// 	return arr[re
