// // find the first and last position of the target in the sorted array...
// Facebook quesition, that much easy
// now we more optimized this code like if the both index is not found then it return [-1, -1]

package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 7, 7, 7, 7, 8, 8, 9, 9, 9, 10}
	tar := 7
	res := search(arr, tar)
	if res[0] != -1 {
		fmt.Printf("The first and last index of the target number is : [%d, %d]", res[0], res[1])
	} else {
		fmt.Println("the first and last pos. of the target number are not found")
	}
}

func search(arr []int, tar int) []int {
	first := findFirst(arr, tar)
	last := findLast(arr, tar)

	if first == -1 && last == -1 {
		return []int{-1, -1}
	}

	return []int{first, last}
}

func findFirst(arr []int, tar int) int {
	start, end := 0, len(arr)-1
	res := -1
	for start <= end {
		mid := start + (end-start)/2
		if tar == arr[mid] {
			res = mid
			end = mid - 1
		} else if tar < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return res
}

func findLast(arr []int, tar int) int {
	start, end := 0, len(arr)-1
	res := -1
	for start <= end {
		mid := start + (end-start)/2
		if tar == arr[mid] {
			res = mid
			start = mid + 1
		} else if tar < arr[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return res
}

// func search(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}

// 	return []int{first, last}
// }

// func findFirst(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			end = mid - 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			start = mid + 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 7, 9, 9, 10}
// 	tar := 7
// 	res := search(arr, tar)
// 	if res[0] != -1 {
// 		fmt.Printf("The first and the last index of the target is : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Taget number is not found")
// 	}
// }

// func search(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)
// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}
// 	return []int{first, last}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 7, 9, 9, 10}
// 	tar := 7
// 	res := search(arr, tar)
// 	if res[0] != -1 {
// 		fmt.Printf("The first and last position of the target number index is : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}
// 	return []int{first, last}
// }

// func findFirst(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			end = mid - 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			start = mid + 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 7, 9, 9, 10}
// 	tar := 7
// 	res := search(arr, tar)
// 	if res != nil {
// 		fmt.Printf("zthe first and the last position of the target number index is : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == 0 && last == 0 {
// 		return []int{-1, -1}
// 	}
// 	return []int{first, last}
// }

// func findFirst(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			end = mid - 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if arr[mid] == tar {
// 			res = mid
// 			start = mid + 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 7, 7, 7, 7, 9, 9, 10}
// 	tar := 7
// 	res := search(arr, tar)
// 	if res != nil {
// 		fmt.Printf("The first and the last position of the target numer index is : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func search(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}
// 	return []int{first, last}
// }

// func findFirst(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			end = mid - 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			start = mid + 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
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
// 	if res != nil {
// 		fmt.Printf("The first and the last position of the target number index is : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
// }

// func firstAndLast(arr []int, tar int) []int {
// 	first := findFirst(arr, tar)
// 	last := findLast(arr, tar)

// 	if first == -1 && last == -1 {
// 		return []int{-1, -1}
// 	}
// 	return []int{first, last}
// }

// func findFirst(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			end = mid - 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, tar int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if tar == arr[mid] {
// 			res = mid
// 			start = mid + 1
// 		} else if tar < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
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
// 	if res != nil {
// 		fmt.Printf("THe first and last position of the target number of index : [%d, %d]\n", res[0], res[1])
// 	} else {
// 		fmt.Println("Target not found")
// 	}
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
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if x == arr[mid] {
// 			res = mid
// 			end = mid - 1
// 		} else if x < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

// func findLast(arr []int, x int) int {
// 	start, end := 0, len(arr)-1
// 	res := -1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if x == arr[mid] {
// 			res = mid
// 			start = mid + 1
// 		} else if x < arr[mid] {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return res
// }

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
