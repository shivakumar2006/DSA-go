package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 1, 4, 3, 2, 6}
	res := SelectionSort(arr)
	fmt.Println(res)
}

// func SelectionSort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[j] < arr[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		if minIndex != i {
// 			temp := arr[i]
// 			arr[i] = arr[minIndex]
// 			arr[minIndex] = temp
// 		}
// 	}
// 	return arr
// }

// func SelectionSort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[j] < arr[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		if minIndex != i {
// 			temp := arr[i]
// 			arr[i] = arr[minIndex]
// 			arr[minIndex] = temp
// 		}
// 	}
// 	return arr
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 1, 4, 2, 6}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[j] < arr[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		if minIndex != i {
// 			temp := arr[i]
// 			arr[i] = arr[minIndex]
// 			arr[minIndex] = temp
// 		}
// 	}
// 	return arr
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 1, 4, 2, 6}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[j] < arr[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		if minIndex != i {
// 			temp := arr[i]
// 			arr[i] = arr[minIndex]
// 			arr[minIndex] = temp
// 		}
// 	}
// 	return arr
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{5, 1, 4, 2, 6}
// 	res := sort(arr)
// 	fmt.Println(res)
// }

// func sort(arr []int) []int {
// 	for i := 0; i < len(arr); i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[j] < arr[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		if minIndex != i {
// 			temp := arr[i]
// 			arr[i] = arr[minIndex]
// 			arr[minIndex] = temp
// 		}
// 	}
// 	return arr
// }
