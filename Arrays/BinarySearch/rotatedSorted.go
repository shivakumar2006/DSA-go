// search in rotated sorted array

package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
	tar := 1
	res := search(arr, tar)
	if res != 0 {
		fmt.Printf("The target are found at the index : %d\n", res)
	} else {
		fmt.Println("not found")
	}
}

func search(arr []int, tar int) int {
	pivot := findPivot(arr)

	if pivot == -1 {
		return binarySearch(arr, tar, 0, len(arr)-1)
	}

	if tar == arr[pivot] {
		return pivot
	}

	if tar >= arr[0] {
		return binarySearch(arr, tar, 0, pivot)
	}
	return binarySearch(arr, tar, pivot+1, len(arr)-1)
}

func findPivot(arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		mid := start + (end-start)/2
		if mid < end && arr[mid] > arr[mid+1] {
			return mid
		} else if mid > start && arr[mid] < arr[mid-1] {
			return mid - 1
		} else if arr[mid] >= arr[start] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func binarySearch(arr []int, tar int, start, end int) int {
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
// 		} else if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 1
// 	res := search(arr, tar)
// 	if res != 0 {
// 		fmt.Printf("The target are found at the index : %d\n", res)
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
// 		} else if arr[mid] <= arr[start] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }

// func binarySearch(arr []int, tar int, start, end int) int {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 1
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("Þhe target element is found at the index : %d\n", res)
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
// 	for start < end {
// 		mid := start + (end-start)/2
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		} else if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		} else if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 1
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element index is found : %d\n", res)
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
// 		} else if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		} else if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 1
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("Th target number index is : %d\n", res)
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
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 1
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target number index is : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	} else if tar == arr[pivot] {
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
// 		} else if mid > start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		} else if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 2
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element index : %d\n", res)
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
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 1
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element index is : %d\n", res)
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
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 0
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element index is : %d\n", res)
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
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 6
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The taret element index is : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}
// 	if arr[pivot] == tar {
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
// 		if arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	tar := 6
// 	res := search(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The target element index is : %d\n", res)
// 	} else {
// 		fmt.Println("Taret not found")
// 	}
// }

// func search(arr []int, tar int) int {
// 	pivot := findPivot(arr)

// 	if pivot == -1 {
// 		return binarySearch(arr, tar, 0, len(arr)-1)
// 	}

// 	if arr[pivot] == tar {
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
// 		if mid >= start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println("The pivot is found at the index : ", findPivot(arr))
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2

// 		// if pivot is lies on the mid then return mid and also mid is < end
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		}
// 		// if mid of the array is less than mid - 1 then return mid - 1th index
// 		if mid >= start && arr[mid] < arr[mid-1] {
// 			return mid - 1
// 		}

// 		// now find where need to go
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println("THe pivot is found at the index : ", findPivot(arr))
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
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println("The pivot is found at the index : ", findPivot(arr))
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
// 		if arr[mid] <= arr[start] {
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
// 	arr := []int{3, 4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println("The pivot is found at the index : ", findPivot(arr))
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2

// 		//check if mid is pivot
// 		if mid < end && arr[mid] > arr[mid+1] {
// 			return mid
// 		} else if mid > start && arr[mid] < arr[mid-1] { // if mid - 1 is pivot then
// 			return mid - 1
// 		} else if arr[mid] <= arr[start] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }
