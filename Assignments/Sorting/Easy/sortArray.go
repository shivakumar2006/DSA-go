// Sort array by parity
// https://leetcode.com/problems/sort-array-by-parity/description/

package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	evenSet := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			evenSet = append(evenSet, num)
		}
	}

	oddSet := []int{}
	for _, num := range nums {
		if num%2 != 0 {
			oddSet = append(oddSet, num)
		}
	}

	nums = append(evenSet, oddSet...)
	return nums
}

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}
