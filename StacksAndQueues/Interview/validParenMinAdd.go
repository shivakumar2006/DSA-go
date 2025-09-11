// Minimum add to make valid parentheses...
// Meta question

package main

import "fmt"

func minAddToMakeValid(s string) int {
	stack := []rune{}

	for _, ch := range s {
		if ch == ')' {
			// if stack top is '(', pop it.
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				// otherwise push ')'
				stack = append(stack, ch)
			}
		} else {
			// push '('
			stack = append(stack, ch)
		}
	}
	return len(stack)
}

func main() {
	fmt.Println(minAddToMakeValid("())"))    // 1
	fmt.Println(minAddToMakeValid("((("))    // 3
	fmt.Println(minAddToMakeValid("()"))     // 0
	fmt.Println(minAddToMakeValid("()))((")) // 4
}
