// Largest Number
// https://leetcode.com/problems/largest-number/description/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	// convert all numbers into string
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	// sort strings with custom comparisons
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	// if the largest number is 0, return 0
	if strs[0] == "0" {
		return "0"
	}

	//concatinate all strings
	return strings.Join(strs, "")
}

func main() {
	nums := []int{10, 2}
	fmt.Println(largestNumber(nums))
}
