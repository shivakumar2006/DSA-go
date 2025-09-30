// check if N and it's double is exists
// https://leetcode.com/problems/check-if-n-and-its-double-exist/

package main

import "fmt"

func checkIfExist(arr []int) bool {
	set := make(map[int]bool)

	for _, x := range arr {
		if set[2*x] || (x%2 == 0 && set[x/2]) {
			return true
		}
		set[x] = true
	}
	return false
}

func main() {
	arr := []int{10, 2, 5, 3}
	fmt.Println(checkIfExist(arr))
}
