// Third MaxNum
// https://leetcode.com/problems/third-maximum-number/description/

package main

import "fmt"

func thirdMax(nums []int) int {
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	// convert map to slice
	distinctNum := []int{}
	for num := range set {
		distinctNum = append(distinctNum, num)
	}

	// sort elements in descending order
	for i := 1; i < len(distinctNum); i++ {
		j := i
		for j > 0 && distinctNum[j] > distinctNum[j-1] {
			swap(distinctNum, j, j-1)
			j--
		}
	}

	// return third max if it's exist otherwise max
	if len(distinctNum) >= 3 {
		return distinctNum[2]
	}
	return distinctNum[0]
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{2, 2, 3, 1}
	fmt.Println(thirdMax(nums))
}
