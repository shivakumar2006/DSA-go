// // Game of two stacks
// // google and amazon level question

package main

import "fmt"

func gameOfTwoStacks(limit int, stack1 []int, stack2 []int) (int, []int) {
	return solve(limit, stack1, stack2, 0, 0, []int{})
}

func solve(limit int, stack1, stack2 []int, sum int, count int, picked []int) (int, []int) {
	if sum > limit {
		return -999999, []int{}
	}

	bestCount := count
	bestPicked := append([]int{}, picked...) // copy slice

	// take from stack 1
	if len(stack1) > 0 {
		c1, p1 := solve(limit, stack1[1:], stack2, sum+stack1[0], count+1, append(picked, stack1[0]))
		if c1 > bestCount {
			bestCount = c1
			bestPicked = p1
		}
	}

	// take from stack 2
	if len(stack2) > 0 {
		c2, p2 := solve(limit, stack1, stack2[1:], sum+stack2[0], count+1, append(picked, stack2[0]))
		if c2 > bestCount {
			bestCount = c2
			bestPicked = p2
		}
	}
	return bestCount, bestPicked
}

func main() {
	stack1 := []int{4, 2, 4, 6, 1}
	stack2 := []int{2, 1, 9, 5}
	limit := 10

	count, removed := gameOfTwoStacks(limit, stack1, stack2)
	fmt.Println("Max elements removes : ", count)
	fmt.Println("Removed element : ", removed)
}

// package main

// import "fmt"

// // function of solve game of two stacks...
// func gameOfTwoStacks(limit int, stack1 []int, stack2 []int) int {
// 	return solve(limit, stack1, stack2, 0, 0)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func solve(limit int, stack1 []int, stack2 []int, sum int, count int) int {
// 	if sum > limit {
// 		return -999999
// 	}

// 	best := count

// 	if len(stack1) > 0 {
// 		best = max(best, solve(limit, stack1[1:], stack2, sum+stack1[0], count+1))
// 	}

// 	if len(stack2) > 0 {
// 		best = max(best, solve(limit, stack1, stack2[1:], sum+stack2[0], count+1))
// 	}

// 	return best
// }

// func main() {
// 	stack1 := []int{4, 2, 4, 6, 1}
// 	stack2 := []int{2, 1, 9, 5}
// 	limit := 10

// 	fmt.Println("Max element removed: ", gameOfTwoStacks(limit, stack1, stack2))
// }
