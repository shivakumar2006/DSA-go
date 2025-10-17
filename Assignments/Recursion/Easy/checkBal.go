// check balanced parenthesis
// https://www.geeksforgeeks.org/dsa/check-for-balanced-parenthesis-without-using-stack/

package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != bracketMap[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	exp1 := "[()]{}{[()()]()}"
	exp2 := "[(])"

	fmt.Println(isValid(exp1)) // true
	fmt.Println(isValid(exp2)) // false
}
