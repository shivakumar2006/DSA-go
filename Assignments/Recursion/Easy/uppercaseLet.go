// First uppercase letter in a string (Iterative and Recursive)
// https://www.geeksforgeeks.org/dsa/first-uppercase-letter-in-a-string-iterative-and-recursive/

package main

import "fmt"

// iterative
// func uppercase(str string) string {
// 	for _, char := range str {
// 		if char >= 'A' && char <= 'Z' {
// 			return string(char)
// 		}
// 	}
// 	return ""
// }

// recursive
func uppercase(str string) string {
	if len(str) == 0 {
		return ""
	}

	if str[0] >= 'A' && str[0] <= 'Z' {
		return string(str[0])
	}

	return uppercase(str[1:])
}

func main() {
	str := "geeksforgeeKs"
	fmt.Println(uppercase(str))
}
