//  Find Kth Bit in Nth Binary String
// http://leetcode.com/problems/find-kth-bit-in-nth-binary-string/description/

package main

import "fmt"

func findKthBit(n, k int) byte {
	if n == 1 {
		return '0'
	}

	length := (1 << n) - 1 // 2 ^ n - 1
	mid := (length / 2) + 1

	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		// right half mirrored
		mirrored := length - k + 1
		bit := findKthBit(n-1, mirrored)
		// invert
		if bit == '0' {
			return '1'
		} else {
			return '0'
		}
	}
}

func main() {
	n := 4
	k := 11
	fmt.Printf("%c\n", findKthBit(n, k))
}
