// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{3, 1, 4, 5, 2}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr)-1; i++ {
// 		correct := arr[i] - 1
// 		if arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, i int, correct int) {
// 	temp := arr[i]
// 	arr[i] = arr[correct]
// 	arr[correct] = temp
// }
