// shuffle String
// https://leetcode.com/problems/shuffle-string/description/

package main

import "fmt"

func restoreString(s string, indices []int) string {
	result := make([]rune, len(s))

	for i, ch := range s {
		result[indices[i]] = ch
	}

	return string(result)
}

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
}
