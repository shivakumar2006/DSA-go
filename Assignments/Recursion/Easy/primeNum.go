// Recursive program for prime number
// https://www.geeksforgeeks.org/dsa/recursive-program-prime-number/

package main

import "fmt"

func isPrime(n int, i int) bool {
	if n <= 2 {
		if n == 2 {
			return true
		}
		return false
	}
	if n%i == 0 {
		return false
	}
	if i*i > n {
		return true
	}
	return isPrime(n, i+1)
}

func main() {
	n := 11
	if isPrime(n, 2) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
