// N Digit numbers with digits in increasing order
// https://www.geeksforgeeks.org/problems/n-digit-numbers-with-digits-in-increasing-order5903/1

package main

import "fmt"

func printIncreasingOrder(n int) {
	var result []int

	var helper func(current, lastDigit, digitsLeft int)
	helper = func(current, lastDigit, digitsLeft int) {
		if digitsLeft == 0 {
			result = append(result, current)
			return
		}
		for d := lastDigit + 1; d <= 9; d++ {
			next := current*10 + d
			helper(next, d, digitsLeft-1)
		}
	}

	if n == 1 {
		fmt.Println("0 ")
	}

	for start := 1; start <= 9; start++ {
		helper(start, start, n-1)
	}

	for _, num := range result {
		fmt.Println(num, " ")
	}
	fmt.Print()
}

func main() {
	n := 1
	printIncreasingOrder(n)
}
