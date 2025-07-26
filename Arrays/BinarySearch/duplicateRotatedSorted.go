// duplicate routated sorted array...

package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 2, 9, 2, 2}
	tar := 9
	res := search(arr, tar)
	if res != -1 {
		fmt.Printf("THe target element found at the index : %d\n", res)
	} else {
		fmt.Println("not found")
	}
}

func search(arr []int, tar int) int {
	pivot := findPivot(arr)

	if pivot == 0 {
		return binarySearch(arr, tar, 0, len(arr)-1)
	}
	if tar == arr[pivot] {
		return pivot
	}
	if tar > arr[0] {
		return binarySearch(arr, tar, 0, pivot)
	}
	return binarySearch(arr, tar, pivot+1, len(arr)-1)
}

func findPivot(arr []int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if mid < end && arr[mid] > arr[mid+1] {
			return mid + 1
		} else if mid > start && arr[mid] < arr[mid-1] {
			return mid - 1
		}

		// handle duplicate
		if arr[start] == arr[mid] && arr[mid] == arr[end] {
			if start < end && arr[start] < arr[start+1] {
				return start + 1
			}
			start++
			if start < end && arr[end] < arr[end-1] {
				return end - 1
			}
			end--
		} else if arr[start] < arr[end] || arr[start] == arr[mid] && arr[start] < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func binarySearch(arr []int, tar, start, end int) int {
	for start <= end {
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
// 	arr := []int{2, 2, 9, 2, 2}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != 0 {
// 		fmt.Printf("The target element is found at the index %d\n", res)
// 	} else {
// 		fmt.Println("not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar == arr[pivot] {
// 		return pivot
// 	}
// 	if tar <= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		} else if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] < arr[start+1] {
// 				return start + 1
// 			}
// 			start++
// 			if start < end && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[end] || arr[start] == arr[mid] && arr[start] < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
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

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar == arr[pivot] {
// 		return pivot
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		} else if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] < arr[start+1] {
// 				return start + 1
// 			}
// 			start++
// 			if start < end && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[end] || arr[start] == arr[mid] && arr[start] < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 9, 2, 2}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("THe target element id found at the index : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar == arr[pivot] {
// 		return pivot
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// handle Duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] < arr[start+1] {
// 				return start + 1
// 			}
// 			start++
// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[end] || arr[start] == arr[mid] && arr[start] > arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 9, 2, 2}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element is found ath the indx : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar == arr[pivot] {
// 		return pivot
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		//Handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] < arr[start+1] {
// 				return start
// 			}
// 			start++
// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[end] || arr[start] == arr[mid] && arr[mid] > arr[end] {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 9, 2, 2}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("THe target element is found at the index : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar == arr[pivot] {
// 		return pivot
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] > arr[start+1] {
// 				return start
// 			}
// 			start++
// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[end] || arr[start] == arr[mid] && arr[mid] > arr[end] {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 9, 2, 2}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element is found at the index : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar == arr[pivot] {
// 		return pivot
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] > arr[start+1] {
// 				return start
// 			}
// 			start++
// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[end] || (arr[start] == arr[mid] && arr[mid] > arr[end]) {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 9, 2, 2, 2}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element is found at the index : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar == arr[pivot] {
// 		return pivot
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] > arr[start+1] {
// 				return start
// 			}
// 			start++

// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[mid] || (arr[start] == arr[mid] && arr[mid] > arr[end]) {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 9, 2, 2, 2}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element index found : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar == arr[pivot] {
// 		return pivot
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// Handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] > arr[start+1] {
// 				return start
// 			}
// 			start++

// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 		} else if arr[start] < arr[mid] || (arr[start] == arr[mid] && arr[mid] > arr[end]) {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 2, 2, 2, 9}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("THe target element indexsound : %d\n", res)
// 	} else {
// 		fmt.Println("Target  not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot != -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// now we handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] > arr[start+1] {
// 				return start
// 			}
// 			start++

// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[mid] || (arr[start] == arr[mid] && arr[mid] > arr[end]) {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 2, 2, 2, 9}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("THe target element indexsound : %d\n", res)
// 	} else {
// 		fmt.Println("Target  not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// now we handle duplicate
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] > arr[start+1] {
// 				return start
// 			}
// 			start++

// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[mid] || (arr[start] == arr[mid] && arr[mid] > arr[end]) {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
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
// 	arr := []int{2, 2, 2, 2, 9}
// 	tar := 9
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("THe target element index found : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}

// 	if tar >= arr[0] {
// 		return binarySearch(arr, tar, 0, pivot-1)
// 	}
// 	return binarySearch(arr, tar, pivot+1, len(arr)-1)
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// Handle Duplicates...
// 		if arr[start] == arr[mid] && arr[mid] == arr[end] {
// 			if start < end && arr[start] > arr[start+1] {
// 				return start
// 			}
// 			start++

// 			if end > start && arr[end] < arr[end-1] {
// 				return end - 1
// 			}
// 			end--
// 		} else if arr[start] < arr[mid] || (arr[start] == arr[mid] && arr[mid] > arr[end]) {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return -1
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
