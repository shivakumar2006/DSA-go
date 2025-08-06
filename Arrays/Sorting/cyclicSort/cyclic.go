package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 1, 4, 5, 2}
	res := sort(arr)
	fmt.Println(res)
}

func sort(arr []int) []int {
	i := 0
	for i < len(arr) {
		correct := arr[i] - 1
		if arr[i] != arr[correct] {
			swap(arr, i, correct)
		} else {
			i++
		}
	}
	return arr
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1
// 		if arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, i int, j int) {
// 	arr[i], arr[j] = arr[j], arr[i]
// }

// func sort(arr []int) []int {
// 	i := 0
// 	for i < len(arr) {
// 		correct := arr[i] - 1 // Our correct index on the arr[i]-1 index
// 		if arr[i] != arr[correct] {
// 			swap(arr, i, correct)
// 		} else {
// 			i++
// 		}
// 	}
// 	return arr
// }

// func swap(arr []int, i int, correct int) {
// 	arr[i], arr[correct] = arr[correct], arr[i]
// }

// This is wrong......

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
