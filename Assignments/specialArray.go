// Special Array With X Elements Greater Than or Equal X
// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 5}
	fmt.Println("number of values greater than or equals to : ", specialArray(arr))
}

func specialArray(arr []int) int {
	sort.Ints(arr)
	n := len(arr)

	for i := 0; i <= n; i++ {
		count := 0
		for _, num := range arr {
			if num >= i {
				count++
			}
		}
		if count == i {
			return i
		}
	}
	return -1
}
