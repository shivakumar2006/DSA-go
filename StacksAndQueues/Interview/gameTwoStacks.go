// Game of two stacks
// google and amazon level question

package main

import "fmt"

// function of solve game of two stacks...
func gameOfTwoStacks(limit int, stack1 []int, stack2 []int) int {
	return solve(limit, stack1, stack2, 0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(limit int, stack1 []int, stack2 []int, sum int, count int) int {
	if sum > limit {
		return -999999
	}

	best := count

	if len(stack1) > 0 {
		best = max(best, solve(limit, stack1[1:], stack2, sum+stack1[0], count+1))
	}

	if len(stack2) > 0 {
		best = max(best, solve(limit, stack1, stack2[1:], sum+stack2[0], count+1))
	}

	return best
}

func main() {
	stack1 := []int{4, 2, 4, 6, 1}
	stack2 := []int{2, 1, 9, 5}
	limit := 10

	fmt.Println("Max element removed: ", gameOfTwoStacks(limit, stack1, stack2))
}
