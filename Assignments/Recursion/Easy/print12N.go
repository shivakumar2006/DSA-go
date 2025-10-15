// Print 1 To N Without Loop
// http://geeksforgeeks.org/problems/print-1-to-n-without-using-loops-1587115620/1

package main

import "fmt"

func print1ToN(n int) {
	if n == 0 {
		return
	}

	print1ToN(n - 1)
	fmt.Print(n, " ")
}

func main() {
	print1ToN(10)
}
