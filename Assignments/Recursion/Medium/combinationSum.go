// combination sum
// https://leetcode.com/problems/combination-sum/description/

package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	backtrack(candidates, target, []int{}, 0, &result)
	return result
}

func backtrack(candidates []int, tar int, current []int, start int, result *[][]int) {
	if tar == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	if tar < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		backtrack(candidates, tar-candidates[i], current, i, result)
		current = current[:len(current)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
