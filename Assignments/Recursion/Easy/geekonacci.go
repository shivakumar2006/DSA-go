// Geek - onacci
// https://www.geeksforgeeks.org/problems/geek-onacci-number/0
// geek(n)=geek(n−1)+geek(n−2)+geek(n−3)

package main

import "fmt"

func geekOnacci(A, B, C, N int) int {
	if N == 1 {
		return A
	}
	if N == 2 {
		return B
	}
	if N == 3 {
		return C
	}
	return geekOnacci(A, B, C, N-1) + geekOnacci(A, B, C, N-2) + geekOnacci(A, B, C, N-3)
}

func main() {
	A, B, C, N := 1, 3, 2, 6
	fmt.Println(geekOnacci(A, B, C, N))
}
