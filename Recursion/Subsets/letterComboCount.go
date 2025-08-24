package main

import "fmt"

func main() {
	ans := pad("", "12")
	fmt.Println(ans)
}

func pad(process, unprocess string) int {
	if len(unprocess) == 0 {
		return 1
	}

	count := 0
	digit := int(unprocess[0] - '0')
	start := (digit - 1) * 3
	end := digit * 3

	for i := start; i < end; i++ {
		ch := string('a' + rune(i))
		count += pad(process+ch, unprocess[1:])
	}
	return count
}
