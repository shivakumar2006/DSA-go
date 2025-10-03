// Count negative numbers in a sorted matrix
// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/description/

package main

import "fmt"

func main() {
	grid := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	fmt.Println("total negative numbers in a gird is : ", countNegatives(grid))
}

func countNegatives(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	row := 0
	col := cols - 1
	for row < rows && col >= 0 {
		if grid[row][col] < 0 {
			count += rows - row
			col--
		} else {
			row++
		}
	}
	return count
}

// func countNegatives(grid [][]int) int {
// 	rows := len(grid)
// 	cols := len(grid[0])
// 	count := 0

// 	row := 0
// 	col := cols - 1 // starts from top rght corner
// 	for row < rows && col >= 0 {
// 		if grid[row][col] < 0 {
// 			count += rows - row
// 			col--
// 		} else {
// 			row++
// 		}
// 	}
// 	return count
// }
