package main

import (
	"fmt"
)

func main() {
	str := []string{"c", "f", "g"}
	tar := "d"
	res := ceiling(str, tar)
	if res != -1 {
		fmt.Printf("THe index of the ceiling number is found : %d\n", res)
	} else {
		fmt.Println("Target index is not found")
	}
}

func ceiling(str []string, x string) int {
	start, end := 0, len(str)-1
	for start <= end {
		mid := start + (end-start)/2
		if x == str[mid] {
			return mid
		} else if x < str[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	if start < len(str) {
		return start
	}
	return -1
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	str := []string{"c", "f", "g"}
// 	tar := "d"
// 	res := ceiling(str, tar)
// 	fmt.Printf("The Index of the ceiling number is found : %d\n", res)
// }

// func ceiling(str []string, x string) int {
// 	l := 0
// 	r := len(str) - 1
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if x == str[mid] {
// 			return mid
// 		} else if x < str[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	if l < len(str) {
// 		return l
// 	}
// 	return -1
// }
