// Height Checker
// https://leetcode.com/problems/height-checker/description/

package main

import "fmt"

func heightChecker(arr []int) int {
	expected := append([]int{}, arr...)

	for i := 0; i < len(expected); i++ {
		for j := 1; j < len(expected)-i; j++ {
			if expected[j] < expected[j-1] {
				swap(expected, j, j-1)
			}
		}
	}

	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			count++
		}
	}

	return count
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	arr := []int{1, 1, 4, 2, 1, 3}
	fmt.Println(heightChecker(arr))
}
