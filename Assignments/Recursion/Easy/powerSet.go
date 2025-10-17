// Power Set in Lexicographic order
// https://www.geeksforgeeks.org/dsa/powet-set-lexicographic-order/

package main

import "fmt"

func generateSubsequences(current, remaining string, result *[]string) {
	if len(remaining) == 0 {
		if len(current) > 0 {
			*result = append(*result, current)
		}
		return
	}
	// exclude first char
	generateSubsequences(current, remaining[1:], result)
	// include first char
	generateSubsequences(current+string(remaining[0]), remaining[1:], result)
}

func main() {
	s := "abc"
	result := []string{}
	generateSubsequences("", s, &result)
	fmt.Println(result)
}
