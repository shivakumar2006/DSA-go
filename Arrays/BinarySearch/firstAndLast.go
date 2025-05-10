// // find the first and last position of the target in the sorted array...
// Favebook quesition, that much easy
// now we more optimized this code like if the both index is not found then it return [-1, -1]

package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 7, 7, 8, 8, 8, 8, 10}
	tar := 8
	res := firstAndLast(arr, tar)
	fmt.Printf("The first and last position of the target index is : [%d, %d]\n", res[0], res[1])
}

func firstAndLast(arr []int, tar int) []int {
	first := findFirst(arr, tar)
	last := findLast(arr, tar)

	if first == -1 && last == -1 {
		return []int{-1, -1}
	}

	return []int{first, last}
}

func findFirst(arr []int, x int) int {
	l, r := 0, len(arr)-1
	res := -1
	for l <= r {
		mid := (l + r) / 2
		if x == arr[mid] {
			res = mid
			r = mid - 1
		} else if x > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}

func findLast(arr []int, x int) int {
	l, r := 0, len(arr)-1
	res := -1
	for l <= r {
		mid := (l + r) / 2
		if x == arr[mid] {
			res = mid
			l = mid + 1
		} else if x < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 8, 8, 8, 8, 10}
// 	tar := 8
// 	res := firstAndLast(arr, tar)
// 	fmt.Printf("The first and last position of the target index is : [%d, %d]\n", res[0], res[1])
// }

// func firstAndLast(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}
// 	return []int{first, last}
// }

// func findFirst(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	res := 0
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if arr[mid] == x {
// 			res = mid
// 			r = mid - 1
// 		} else if x > arr[mid] {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	res := 0
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if arr[mid] == x {
// 			res = mid
// 			l = mid + 1
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 7, 8, 8, 10}
// 	tar := 7
// 	res := firstAndLast(arr, tar)
// 	fmt.Printf("The fist and last position of the target index is : [%d, %d]\n", res[0], res[1])
// }

// func firstAndLast(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}

// 	return []int{first, last}
// }

// func findFirst(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			r = mid - 1
// 		} else if x > arr[mid] {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			l = mid + 1
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 7, 8, 8, 10}
// 	tar := 7
// 	res := firstAndLast(arr, tar)
// 	fmt.Printf("The first and last posotion of the target index is : [%d, %d]\n", res[0], res[1])
// }

// func firstAndLast(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}

// 	return []int{first, last}
// }

// func findFirst(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			r = mid - 1
// 		} else if x > arr[mid] {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			l = mid + 1
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 8, 8, 8, 8, 9, 9, 10}
// 	tar := 8
// 	res := firstAndLast(arr, tar)
// 	fmt.Printf("The first and last position of the target is : [%d, %d]\n", res[0], res[1])
// }

// func firstAndLast(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}

// 	return []int{first, last}
// }

// func findFirst(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			r = mid - 1
// 		} else if x > arr[mid] {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, x int) int {
// 	l, r := 0, len(arr)-1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			l = mid + 1
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 7, 8, 8, 10}
// 	tar := 8
// 	res := firstAndLast(arr, tar)
// 	fmt.Printf("The first and last position of the target is : [%d, %d]\n", res[0], res[1])
// }

// func firstAndLast(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}
// 	return []int{first, last}
// }

// func findFirst(arr []int, x int) int {
// 	l := 0
// 	r := len(arr) - 1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			r = mid - 1
// 		} else if x > arr[mid] {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, x int) int {
// 	l := 0
// 	r := len(arr) - 1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			l = mid + 1
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 7, 8, 8, 10}
// 	tar := 7
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)
// 	if first != -1 && last != -1 {
// 		fmt.Printf("The first and last position of the target index is : [%d, %d]\n", first, last)
// 	} else {
// 		fmt.Println("The target both position is not found...")
// 	}
// }

// func findFirst(arr []int, x int) int {
// 	l := 0
// 	r := len(arr) - 1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			r = mid - 1
// 		} else if x > arr[mid] {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, x int) int {
// 	l := 0
// 	r := len(arr) - 1
// 	res := -1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == arr[mid] {
// 			res = mid
// 			l = mid + 1
// 		} else if x < arr[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return res
// }
