// valid perfect square
// https://leetcode.com/problems/valid-perfect-square/description/

package main

import "fmt"

func main() {
	num := 1
	fmt.Println("Is perfect square : ", isPerfectSquare(num))
}

func isPerfectSquare(num int) bool {
	start := 1
	end := num

	for start <= end {
		mid := start + (end-start)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}
