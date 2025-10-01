// Reach a number
// https://leetcode.com/problems/reach-a-number/

package main

import "fmt"

func main() {
	tar := 2
	fmt.Println(reachNumber(tar))
}

func reachNumber(tar int) int {
	if tar < 0 {
		tar = -tar
	}

	sum, steps := 0, 0
	for sum < tar || (sum-tar)%2 != 0 {
		steps++
		sum += steps
	}
	return steps
}

// package main

// import (
// 	"fmt"
// )

// func reachNumber(target int) int {
// 	if target < 0 {
// 		target = -target
// 	}

// 	sum, steps := 0, 0
// 	for sum < target || (sum-target)%2 != 0 {
// 		steps++
// 		sum += steps
// 	}
// 	return steps
// }

// func main() {
// 	target := 3
// 	fmt.Println(reachNumber(target))
// }
