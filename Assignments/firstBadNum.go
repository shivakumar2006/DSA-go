// first bad numbers
// https://leetcode.com/problems/first-bad-version/description/

package main

import "fmt"

var bad int = 4

func isBadVersion(version int) bool {
	if version == bad {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	start := 1
	end := n

	for start < end {
		mid := start + (end-start)/2
		res := isBadVersion(mid)
		if res == true {
			end = mid
		} else if res == false {
			start = mid + 1
		}
	}
	return start
}

func main() {
	n := 5
	fmt.Println("THe bad number is : ", firstBadVersion(n))
}
