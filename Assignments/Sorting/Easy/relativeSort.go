// Relative Sort Array
// https://leetcode.com/problems/relative-sort-array/description/

package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1, arr2 []int) []int {
	count := make(map[int]int)

	for _, num := range arr1 {
		count[num]++
	}

	result := []int{}
	for _, num := range arr2 {
		for count[num] > 0 {
			result = append(result, num)
			count[num]--
		}
		delete(count, num)
	}

	leftOver := []int{}
	for num, freq := range count {
		for i := 0; i < freq; i++ {
			leftOver = append(leftOver, num)
		}
	}

	sort.Ints(leftOver)

	result = append(result, leftOver...)

	return result
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

// func relativeSortArray(arr1, arr2 []int) []int {
// 	count := make(map[int]int)

// 	for _, num := range arr1 {
// 		count[num]++
// 	}

// 	result := []int{}

// 	for _, num := range arr2 {
// 		for count[num] > 0 {
// 			result = append(result, num)
// 			count[num]--
// 		}
// 		delete(count, num)
// 	}

// 	leftOver := []int{}
// 	for num, freq := range count {
// 		for i := 0; i < freq; i++ {
// 			leftOver = append(leftOver, num)
// 		}
// 	}

// 	sort.Ints(leftOver)

// 	result = append(result, leftOver...)

// 	return result
// }
