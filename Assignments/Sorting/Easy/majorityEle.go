//Majority Element
// https://leetcode.com/problems/majority-element/description/

package main

import "fmt"

func majorityElement(nums []int) int {
	count := make(map[int]int)
	n := len(nums)
	for _, num := range nums {
		count[num]++
		if count[num] > n/2 {
			return num
		}
	}
	return -1
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums))
}
