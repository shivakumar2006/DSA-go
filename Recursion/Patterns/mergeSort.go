package main

import "fmt"

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	sorted := mergeSort(arr)
	fmt.Println(sorted)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// func mergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2
// 	left := mergeSort(arr[:mid])
// 	right := mergeSort(arr[mid:])

// 	return merge(left, right)
// }

// func merge(left, right []int) []int {
// 	result := []int{}
// 	i, j := 0, 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}

// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)

// 	return result
// }

// func mergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2
// 	left := mergeSort(arr[:mid])
// 	right := mergeSort(arr[mid:])

// 	return merge(left, right)
// }

// func merge(left, right []int) []int {
// 	result := []int{}
// 	i, j := 0, 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}

// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)

// 	return result
// }

// func mergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2
// 	left := mergeSort(arr[:mid])
// 	right := mergeSort(arr[mid:])

// 	return merge(left, right)
// }

// func merge(left, right []int) []int {
// 	result := []int{}
// 	i, j := 0, 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}

// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)

// 	return result
// }

// func mergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2
// 	left := mergeSort(arr[:mid])
// 	right := mergeSort(arr[mid:])

// 	return merge(left, right)
// }

// func merge(left []int, right []int) []int {
// 	result := []int{}
// 	i, j := 0, 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}

// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)

// 	return result
// }

// func mergeSort(arr []int) []int {
// 	if len(arr) <= 1 {
// 		return arr
// 	}

// 	mid := len(arr) / 2
// 	left := mergeSort(arr[:mid])
// 	right := mergeSort(arr[mid:])

// 	return merge(left, right)
// }

// func merge(left []int, right []int) []int {
// 	result := []int{}
// 	i, j := 0, 0

// 	for i < len(left) && j < len(right) {
// 		if left[i] < right[j] {
// 			result = append(result, left[i])
// 			i++
// 		} else {
// 			result = append(result, right[j])
// 			j++
// 		}
// 	}

// 	//append remaining elements
// 	result = append(result, left[i:]...)
// 	result = append(result, right[j:]...)

// 	return result
// }
