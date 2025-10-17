// find all palindromic partition of a string
// https://www.geeksforgeeks.org/dsa/given-a-string-print-all-possible-palindromic-partition/

package main

import "fmt"

func palindrom(s string) [][]string {
	res := [][]string{}
	palindromeHelper(s, []string{}, &res)
	return res
}

func palindromeHelper(s string, path []string, res *[][]string) {
	if len(s) == 0 {
		temp := make([]string, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := 1; i <= len(s); i++ {
		substr := s[:i]
		if isPalindrome(substr) {
			palindromeHelper(s[i:], append(path, substr), res) // use s[i:] not s[:i]
		}
	}
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true // must return true here
}

func main() {
	s := "geeks"
	fmt.Println(palindrom(s))
}
