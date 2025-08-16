// sum of numers in recursion...

package main

import "fmt"

func main() {
	ans := fact(5)
	fmt.Println(ans)
}

func fact(n int) int {
	if n < 1 {
		return 0
	}

	return n + fact(n-1)
}
