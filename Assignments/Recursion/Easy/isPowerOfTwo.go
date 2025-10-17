// is power of 2
// https://leetcode.com/problems/power-of-two/description/

package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || n%2 != 0 && n != 1 {
		return false
	}

	return isPowerOfTwo(n / 2)
}

func main() {
	n := 6
	fmt.Println(isPowerOfTwo(n))
}
