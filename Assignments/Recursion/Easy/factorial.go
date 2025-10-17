// factorial
// https://www.hackerrank.com/challenges/30-recursion/problem

package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	n := 3
	fmt.Println(factorial(n))
}
