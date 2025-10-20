// count good num
// https://leetcode.com/problems/count-good-numbers/description/

package main

import "fmt"

const mod int64 = 1e9 + 7

func pow(x, y int64) int64 {
	result := int64(1)
	for y < 0 {
		if y%2 == 1 {
			result = (result * x) % mod
		}
		x = (x * x) % mod
		y /= 2
	}

	return result
}

func countGoodNumbers(n int64) int {
	even := (n + 1) / 2
	odd := n / 2
	return int((pow(5, even) * pow(4, odd)) % mod)
}

func main() {
	n := int64(1)
	fmt.Println(countGoodNumbers(n))
}
