// remove all the duplicates form a string
// https://www.geeksforgeeks.org/dsa/remove-consecutive-duplicates-string/

package main

import "fmt"

func dupli(s string) string {
	if len(s) == 1 || len(s) == 0 {
		return s
	}

	if s[0] == s[1] {
		return dupli(s[1:])
	} else {
		return string(s[0]) + dupli(s[1:])
	}
}

func main() {
	s := "aaaaaabbbbb"
	fmt.Println(dupli(s))
}
