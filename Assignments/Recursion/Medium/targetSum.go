// target sum
// https://leetcode.com/problems/target-sum/description/

package main

import "fmt"

func findTargetSumWays(nums []int, tar int) int {
	count := 0
	dfs(nums, 0, 0, tar, &count)
	return count
}

func dfs(nums []int, i, currentSum, tar int, count *int) {
	if i == len(nums) {
		if currentSum == tar {
			*count++
		}
		return
	}
	// add current number
	dfs(nums, i+1, currentSum+nums[i], tar, count)
	// subtract current number
	dfs(nums, i+1, currentSum-nums[i], tar, count)
}

func main() {
	arr := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(arr, target))
}
