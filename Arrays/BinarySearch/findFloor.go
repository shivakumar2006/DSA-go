package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 4, 5, 6, 9, 11, 14, 15, 20}
	tar := 18
	res := find(arr, tar)
	if res != -1 {
		fmt.Printf("The floor of the target elwement is foudn at the index : %d\n", res)
	} else {
		fmt.Println("Target not found")
	}
}

func find(arr []int, tar int) int {
	start, end := 0, len(arr)-1
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
	if end >= 0 {
		return end
	}
	return -1
}

// func find(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if arr[mid] == tar {
// 			return mid
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	if end >= 0 {
// 		return end
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 11, 14, 15, 20}
// 	tar := 18
// 	res := find(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("THe floor of the target element is found : %d\n", res)
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func find(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
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
// 	if end >= 0 {
// 		return end
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 11, 14, 16, 20}
// 	tar := 18
// 	res := findFloor(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The floor of the target number is found at the index : %d\n", res)
// 	} else {
// 		fmt.Println("The floor of the target is not found")
// 	}
// }

// func findFloor(arr []int, x int) int {
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
// 	if end >= 0 {
// 		return end
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 11, 14, 16, 20}
// 	tar := 10
// 	res := findFloor(arr, tar)
// 	if res != -1 {
// 		fmt.Printf("The floor of the target number is found at the index of : %d\n", res)
// 	} else {
// 		fmt.Println("THe floor of the target numner is not found...")
// 	}
// }

// func findFloor(arr []int, x int) int {
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
// 	if r >= 0 {
// 		return r
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 11, 14, 16, 20}
// 	tar := 10
// 	res := findFloor(arr, tar)
// 	fmt.Printf("The floor of the target number us found at the index of : %d\n", res)
// }

// func findFloor(arr []int, x int) int {
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
// 	if r >= 0 {
// 		return r
// 	}
// 	return -1
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{2, 4, 5, 6, 9, 11, 14, 16, 20}
// 	tar := 2
// 	res := findFloor(arr, tar)
// 	fmt.Printf("The floor of the taret number is found at the index of : %d\n", res)
// }

// func findFloor(arr []int, x int) int {
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
// 	if r >= 0 {
// 		return r
// 	}
// 	return -1
// }
