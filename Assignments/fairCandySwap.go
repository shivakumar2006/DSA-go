// Fair Candy Swap
// https://leetcode.com/problems/fair-candy-swap/

package main

import "fmt"

func main() {
	aliceSize := []int{1, 1}
	bobSize := []int{1, 2}

	fmt.Println(fairCandySwap(aliceSize, bobSize))
}

func fairCandySwap(aliceSize, bobSize []int) []int {
	sumAlice, sumBob := 0, 0
	for _, x := range aliceSize {
		sumAlice += x
	}

	for _, y := range bobSize {
		sumBob += y
	}

	tar := (sumAlice - sumBob) / 2 // 0.5

	bobSet := make(map[int]bool)
	for _, y := range bobSize {
		bobSet[y] = true
	}

	for _, x := range aliceSize {
		y := x - tar
		if bobSet[y] {
			return []int{x, y}
		}
	}

	return []int{}
}
