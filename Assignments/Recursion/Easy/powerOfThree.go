// power of 3
// https://leetcode.com/problems/power-of-three/description/

package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || n%3 != 0 {
		return false
	}

	return isPowerOfThree(n / 3)
}

func main() {
	n := 0
	fmt.Println(isPowerOfThree(n))
}
