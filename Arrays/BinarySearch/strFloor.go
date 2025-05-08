package main

import "fmt"

func main() {
	str := []string{"c", "e", "h"}
	tar := "g"
	res := floor(str, tar)
	fmt.Printf("The floor of the target character is found at the index of : %d\n", res)
}

func floor(str []string, x string) int {
	l := 0
	r := len(str) - 1
	for l <= r {
		mid := (l + r) / 2
		if x == str[mid] {
			return mid
		} else if x < str[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if r >= 0 {
		return r
	}
	return -1
}
