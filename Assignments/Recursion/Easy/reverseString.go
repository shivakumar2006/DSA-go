// REverse string
// https://leetcode.com/problems/reverse-string/description/

package main

import "fmt"

// func reverseString(s []byte) {
// 	left := 0
// 	right := len(s) - 1

// 	for left < right {
// 		s[left], s[right] = s[right], s[left]
// 		left++
// 		right--
// 	}
// }

func reverseString(s []byte) {
	reverseHelper(s, 0, len(s)-1)
}

func reverseHelper(s []byte, left, right int) {
	if left >= right {
		return
	}

	s[left], s[right] = s[right], s[left]

	reverseHelper(s, left+1, right-1)
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))
}
