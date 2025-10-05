// contains Duplicate
// https://leetcode.com/problems/contains-duplicate/description/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	set := make(map[int]int)
	for _, num := range nums {
		set[num]++
		if set[num] > 1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 4}
	fmt.Println(containsDuplicate(nums))
}
