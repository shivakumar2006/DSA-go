// Sum of the digit of a number using recursion
// https://www.geeksforgeeks.org/dsa/sum-digit-number-using-recursion/

package main

import "fmt"

func sumOfDigit(n int) int {
	if n == 0 {
		return 0
	}

	return (n % 10) + sumOfDigit(n/10)
}

func main() {
	input := 12345
	fmt.Println(sumOfDigit(input))
}
