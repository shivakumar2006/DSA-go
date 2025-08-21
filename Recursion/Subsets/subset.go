package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	ans := subsets(arr)
	fmt.Println(ans)
}

func subsets(arr []int) [][]int {
	result := [][]int{{}}

	for _, num := range arr {
		newSubsets := [][]int{}
		for _, subset := range result {
			newSubset := append([]int{}, subset...)
			newSubset = append(newSubset, num)
			newSubsets = append(newSubsets, newSubset)
		}
		result = append(result, newSubsets...)
	}
	return result
}

// func subsets(arr []int) [][]int {
// 	result := [][]int{{}}

// 	for _, num := range arr {
// 		newSubsets := [][]int{}
// 		for _, subset := range result {
// 			newSubset := append([]int{}, subset...)
// 			newSubset = append(newSubset, num)
// 			newSubsets = append(newSubsets, newSubset)
// 		}
// 		result = append(result, newSubsets...)
// 	}
// 	return result
// }

// func subsets(arr []int) [][]int {
// 	result := [][]int{{}}

// 	for _, num := range arr {
// 		newSubsets := [][]int{}
// 		for _, subset := range result {
// 			newSubset := append([]int{}, subset...)
// 			newSubset = append(newSubset, num)
// 			newSubsets = append(newSubsets, newSubset)
// 		}
// 		result = append(result, newSubsets...)
// 	}
// 	return result
// }

// func subsets(arr []int) [][]int {
// 	result := [][]int{{}}

// 	for _, num := range arr {
// 		newSubsets := [][]int{}
// 		for _, subset := range result {
// 			// create a new subsetby adding num to exisiting subset
// 			newSubset := append([]int{}, subset...)
// 			newSubset = append(newSubset, num)
// 			newSubsets = append(newSubsets, newSubset)
// 		}
// 		result = append(result, newSubsets...)
// 	}
// 	return result
// }
