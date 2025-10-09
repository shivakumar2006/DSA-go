// Rank Transform of an array
// https://leetcode.com/problems/rank-transform-of-an-array/description/

package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// make a copy of an array and sort it
	sortedArr := append([]int{}, arr...)
	sort.Ints(sortedArr)

	// assign rank to unique elements
	rankMap := make(map[int]int)
	rank := 1

	for _, num := range sortedArr {
		if _, exists := rankMap[num]; !exists {
			rankMap[num] = rank
			rank++
		}
	}

	// replace elements in arr with their rank
	for i, num := range arr {
		arr[i] = rankMap[num]
	}

	return arr
}

func main() {
	arr := []int{40, 10, 20, 30}
	fmt.Println(arrayRankTransform(arr))
}
