// // Valid parantheses...
// // Google and Meta question

package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if ch == ')' && top != '(' {
				return false
			}
			if ch == '}' && top != '{' {
				return false
			}
			if ch == ']' && top != '[' {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("([{}])"))
}

// package main

// import "fmt"

// func isValid(s string) bool {
// 	stack := []rune{} // use rune slice to store characters

// 	for _, ch := range s {
// 		if ch == '(' || ch == '{' || ch == '[' {
// 			// push opening bracket
// 			stack = append(stack, ch)
// 		} else {
// 			// for closing brackets, check stack
// 			if len(stack) == 0 {
// 				return false
// 			}

// 			// pop top of stack
// 			top := stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]

// 			if ch == ')' && top != '(' {
// 				return false
// 			}
// 			if ch == '}' && top != '{' {
// 				return false
// 			}
// 			if ch == ']' && top != '[' {
// 				return false
// 			}
// 		}
// 	}
// 	// at the end, stack must be empty
// 	return len(stack) == 0
// }

// func main() {
// 	fmt.Println(isValid("()"))     // true
// 	fmt.Println(isValid("()[]{}")) // true
// 	fmt.Println(isValid("(]"))     // false
// 	fmt.Println(isValid("([)]"))   // false
// 	fmt.Println(isValid("{[]}"))   // true
// }
