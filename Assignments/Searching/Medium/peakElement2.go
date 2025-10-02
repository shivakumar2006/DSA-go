// Find a peak element
// https://leetcode.com/problems/find-a-peak-element-ii/description/

package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 4},
		{3, 2},
	}
	fmt.Println(findPeakGrid(arr))
}

func findPeakGrid(arr [][]int) []int {
	row := len(arr)
	col := len(arr[0])
	startCol, endCol := 0, col-1

	for startCol <= endCol {
		midCol := startCol + (endCol-startCol)/2

		// find the row index of the max element in this column
		maxRow := 0
		for r := 0; r < row; r++ {
			if arr[r][midCol] > arr[maxRow][midCol] {
				maxRow = r
			}
		}

		left := -1
		if midCol-1 >= 0 {
			left = arr[maxRow][midCol-1]
		}

		right := -1
		if midCol+1 < col {
			right = arr[maxRow][midCol+1]
		}

		// check if a max element is a peak
		if arr[maxRow][midCol] > left && arr[maxRow][midCol] > right {
			return []int{maxRow, midCol}
		} else if right > arr[maxRow][midCol] {
			startCol = midCol + 1
		} else {
			endCol = midCol - 1
		}
	}
	return []int{-1, -1}
}
