// number of steps to reduce a number to zero
// https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/description/

package main

import "fmt"

func numberOfSteps(nums int) int {
	if nums == 0 {
		return 0
	}

	if nums%2 == 0 {
		return 1 + numberOfSteps(nums/2)
	} else {
		return 1 + numberOfSteps(nums-1)
	}
}

func main() {
	nums := 14
	fmt.Println(numberOfSteps(nums))
}
