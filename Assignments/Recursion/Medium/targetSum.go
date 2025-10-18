// target sum
// https://leetcode.com/problems/target-sum/description/

package main

import "fmt"

func findTargetSumWays(nums []int, tar int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if (tar+total)%2 != 0 || total < abs(tar) {
		return 0
	}

	sum := (tar + total) / 2
	dp := make([]int, sum+1)
	dp[0] = 1

	for _, num := range nums {
		for j := sum; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[sum]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	arr := []int{1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(arr, target))
}
