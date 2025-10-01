// Koko eating banana
// https://leetcode.com/problems/koko-eating-bananas/description/

package main

import (
	"fmt"
	"math"
)

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(minEatingSpeed(piles, h))
}

func minEatingSpeed(piles []int, h int) int {
	start, end := 1, 0

	for _, pile := range piles {
		if pile > end {
			end = pile
		}
	}

	for start <= end {
		mid := start + (end-start)/2
		if canEatAll(piles, h, mid) {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return start
}

func canEatAll(piles []int, h int, mid int) bool {
	hours := 0
	for _, pile := range piles {
		hours += int(math.Ceil(float64(pile) / float64(mid)))
	}
	return hours <= h
}
