// length of string
// https://www.geeksforgeeks.org/dsa/program-for-length-of-a-string-using-recursion/

package main

import "fmt"

func recLen(str string) int {
	if str == "" {
		return 0
	}

	return 1 + recLen(str[1:])
}

func main() {
	str := "abcd"
	fmt.Println(recLen(str))
}
