// Handshakes
// https://www.geeksforgeeks.org/problems/handshakes1303/1

package main

import "fmt"

var memo map[int]int

func count(n int) int {
	if memo == nil {
		memo = make(map[int]int)
	}

	if n <= 2 {
		return 1
	}

	if val, ok := memo[n]; ok {
		return val
	}

	totalWays := 0
	for k := 2; k <= n; k += 2 {
		left := count(k - 2)
		right := count(n - k)
		totalWays += left * right
	}

	memo[n] = totalWays
	return totalWays
}

func main() {
	n := 4
	fmt.Println(count(n))
}
