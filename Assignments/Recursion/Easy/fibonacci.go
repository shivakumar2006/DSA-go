// Fibonacci number
// https://leetcode.com/problems/fibonacci-number/description/

package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	n := 2
	fmt.Print(fib(n))
}
