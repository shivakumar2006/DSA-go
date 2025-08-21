// skip substring

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "xxapplexx"
	fmt.Println(skipSubstring(str, "apple"))
}

func skipSubstring(str string, skip string) string {
	if len(str) == 0 {
		return ""
	}

	// only if current string starts with skip substring
	if strings.HasPrefix(str, skip) {
		// skip the whole substring and continue recursion
		return skipSubstring(str[len(skip):], skip)
	} else {
		// otherwise, take the first character and recurse
		return string(str[0]) + skipSubstring(str[1:], skip)
	}
}
