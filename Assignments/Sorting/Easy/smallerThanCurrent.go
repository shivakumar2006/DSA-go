// How many numbers are smaller than the current number
// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/description/

package main

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	set := append([]int{}, arr...)
	sort.Ints(set)

	setMap := make(map[int]int)
	for i, num := range set {
		if _, exists := setMap[num]; !exists {
			setMap[num] = i
		}
	}

	result := []int{}
	for _, num := range arr {
		result = append(result, setMap[num])
	}

	return result
}

// func smallerNumbersThanCurrent(arr []int) []int {
// 	if len(arr) == 0 {
// 		return []int{}
// 	}

// 	result := []int{}

// 	for i := 0; i < len(arr); i++ {
// 		count := 0
// 		for j := 0; j < len(arr); j++ {
// 			if arr[j] < arr[i] {
// 				count++
// 			}
// 		}
// 		result = append(result, count)
// 	}
// 	return result
// }

func main() {
	arr := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(arr))
}
