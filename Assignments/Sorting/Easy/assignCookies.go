// Assign Cookies
// https://leetcode.com/problems/assign-cookies/description/

package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren(g, s))
}
