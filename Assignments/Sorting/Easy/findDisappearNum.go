// Find All Numbers Disappeared in an Array
// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/

package main

import "fmt"

func findDisappearedNumbers(arr []int) []int {
	set := make(map[int]bool)
	for _, num := range arr {
		set[num] = true
	}

	missing := []int{}
	for i := 1; i <= len(arr); i++ {
		if !set[i] {
			missing = append(missing, i)
		}
	}

	return missing
}

func main() {
	// arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	arr := []int{1, 1}
	fmt.Println(findDisappearedNumbers(arr))
}
