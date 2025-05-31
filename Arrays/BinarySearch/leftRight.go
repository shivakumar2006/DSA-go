package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 0, 1, 0, 1, 0}
	res := split(arr)
	fmt.Println(res)
}

func split(arr []int) []int {
	start, end := 0, len(arr)-1
	res := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			res[start] = 0
			start++
		} else {
			res[end] = 1
			end--
		}
	}
	return res
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{0, 0, 1, 0, 1, 1}
// 	res := leftRight(arr)
// 	fmt.Println(res)
// }

// func leftRight(arr []int) []int {
// 	start, end := 0, len(arr)-1
// 	result := make([]int, len(arr))
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == 0 {
// 			result[start] = 0
// 			start++
// 		} else {
// 			result[end] = 1
// 			end--
// 		}
// 	}
// 	return result
// }
