// pow(x, n)
// https://leetcode.com/problems/powx-n/description/

package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	half := myPow(x, n/2)

	if n%2 == 0 {
		return half * half
	} else {
		return x * half * half
	}
}

func main() {
	fmt.Println(myPow(2, -2))
}
