// skip a character

package main

import "fmt"

func main() {
	str := "bananas"
	result := skipChar(str, 'a')
	fmt.Println(result)
}

func skipChar(s string, skip rune) string {
	if len(s) == 0 {
		return ""
	}

	ch := rune(s[0])
	rest := skipChar(s[1:], skip)

	if ch == skip {
		return rest
	} else {
		return string(ch) + rest
	}
}

// func skipChar(s string, skip rune) string {
// 	// base case
// 	if len(s) == 0 {
// 		return ""
// 	}

// 	ch := rune(s[0])              // current char
// 	rest := skipChar(s[1:], skip) // recursion on the rest of the string

// 	if ch == skip {
// 		return rest // skip this character
// 	} else {
// 		return string(ch) + rest // keep character other than "a" + result of recursion
// 	}
// }
