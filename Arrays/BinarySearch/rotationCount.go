// Rotation count ...

package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(rotation(arr))
}

func rotation(arr []int) int {
	pivot := findPivot(arr)
	return pivot + 1
}

func findPivot(arr []int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if mid < end && arr[mid] > arr[mid+1] {
			return mid
		} else if mid > start && arr[mid] < arr[mid-1] {
			return mid - 1
		} else if arr[mid] <= arr[start] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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
// 			start = mid - 1
// 		}
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
// }

// func findPivot(arr []int) int {
// 	start, end := 0, len(arr)-1
// 	for start < end {
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
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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
// 	arr := []int{4, 5, 6, 7, 0, 1, 2}
// 	fmt.Println(rotation(arr))
// }

// func rotation(arr []int) int {
// 	pivot := findPivot(arr)
// 	return pivot + 1
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
