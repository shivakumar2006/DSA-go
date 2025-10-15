// Product of 2 Numbers using Recursion
// https://www.geeksforgeeks.org/dsa/product-2-numbers-using-recursion/

package main

import "fmt"

func product(x, y int) int {
	if x < y {
		return product(y, x)
	} else if y != 0 {
		return (x + product(x, y-1))
	} else {
		return 0
	}
}

func main() {
	x, y := 5, 2
	fmt.Println(product(x, y))
}
