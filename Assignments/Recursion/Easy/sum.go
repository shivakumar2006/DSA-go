// Sum of First N Natural Numbers Using Recursion
// https://www.geeksforgeeks.org/dsa/sum-of-natural-numbers-using-recursion/

package main

import "fmt"

func sum(n int) int {
	if n == 0 {
		return 0
	}

	return n + sum(n-1)
}

func main() {
	n := 3
	fmt.Println(sum(n))
}
