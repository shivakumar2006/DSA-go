// Search in 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix/description/

package main

import "fmt"

func searchMatrix(arr [][]int, tar int) bool {
	row := 0
	col := len(arr[row]) - 1
	for row < len(arr) && col >= 0 {
		if arr[row][col] == tar {
			return true
		} else if arr[row][col] > tar {
			col--
		} else {
			row++
		}
	}
	return false
}

func main() {
	arr := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{22, 30, 34, 60},
	}
	target := 0
	ans := searchMatrix(arr, target)
	fmt.Println(ans)
}
