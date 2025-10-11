// Group Anagrams
// https://leetcode.com/problems/group-anagrams/description/

package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(str []string) [][]string {
	set := make(map[string][]string)

	for _, word := range str {
		// convert word to slice
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		// sorted version becomes key
		sortedWords := string(runes)

		// added to map
		set[sortedWords] = append(set[sortedWords], word)
	}

	result := [][]string{}
	for _, group := range set {
		result = append(result, group)
	}

	return result
}

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(str))
}
