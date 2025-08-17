// find the given array is sorted or not...

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 7, 44, 33}
	fmt.Println(sorted(arr, 0))
}

func sorted(arr []int, index int) bool {
	if index == len(arr)-1 {
		return false
	}

	return arr[index] < arr[index+1] && sorted(arr, index+1)
}

// func sorted(arr []int, index int) bool {
// 	if index == len(arr)-1 {
// 		return true
// 	}

// 	return arr[index] < arr[index+1] && sorted(arr, index+1)
// }

// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 2, 3, 5, 8, 10}
// 	fmt.Println(sorted(arr, 0))
// }

// func sorted(arr []int, index int) bool {
// 	// base condition
// 	if index == len(arr)-1 {
// 		return true
// 	}

// 	return arr[index] < arr[index+1] && sorted(arr, index+1)
// }
