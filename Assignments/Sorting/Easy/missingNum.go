// Missing Number
// https://leetcode.com/problems/missing-number/description/

package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	total := n * (n + 1) / 2
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return total - sum
}

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}
