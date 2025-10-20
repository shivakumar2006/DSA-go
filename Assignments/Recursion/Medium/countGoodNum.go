// count good num
// https://leetcode.com/problems/count-good-numbers/description/

package main

import "fmt"

func countGoodNumbers(n int64) int {
	// recursive helper
	var helper func(pos int64) int64
	helper = func(pos int64) int64 {
		if pos == n {
			return 1
		}
		if pos%2 == 0 {
			return 5 * helper(pos+1)
		} else {
			return 4 * helper(pos+1)
		}
	}

	return int(helper(0))
}

func main() {
	n := int64(1)
	fmt.Println(countGoodNumbers(n))
}
