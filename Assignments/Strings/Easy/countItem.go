// CountItem matching a rule
// https://leetcode.com/problems/count-items-matching-a-rule/description/

package main

import "fmt"

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	index := 0
	if ruleKey == "color" {
		index = 1
	} else if ruleKey == "name" {
		index = 2
	}

	count := 0
	for _, item := range items {
		if item[index] == ruleValue {
			count++
		}
	}
	return count
}

func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
	ruleKey := "color"
	ruleValue := "silver"
	fmt.Println(countMatches(items, ruleKey, ruleValue))
}
