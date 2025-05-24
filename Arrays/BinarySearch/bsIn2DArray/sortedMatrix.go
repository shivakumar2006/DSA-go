package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	target := 9
	result := search(matrix, target)
	fmt.Println(result) // Output: [2 2]
}

func binarySearch(matrix [][]int, row, start, end, target int) []int {
	for start <= end {
		mid := start + (end-start)/2
		if matrix[row][mid] == target {
			return []int{row, mid}
		}
		if matrix[row][mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return []int{-1, -1}
}

func search(matrix [][]int, target int) []int {
	rows := len(matrix)
	if rows == 0 || len(matrix[0]) == 0 {
		return []int{-1, -1}
	}
	cols := len(matrix[0])

	if rows == 1 {
		return binarySearch(matrix, 0, 0, cols-1, target)
	}

	start := 0
	end := rows - 1
	midCol := cols / 2

	for start < end-1 {
		mid := start + (end-start)/2
		if matrix[mid][midCol] == target {
			return []int{mid, midCol}
		}
		if matrix[mid][midCol] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if matrix[start][midCol] == target {
		return []int{start, midCol}
	}
	if matrix[start+1][midCol] == target {
		return []int{start + 1, midCol}
	}

	if target <= matrix[start][midCol-1] {
		return binarySearch(matrix, start, 0, midCol-1, target)
	}
	if target >= matrix[start][midCol+1] && target <= matrix[start][cols-1] {
		return binarySearch(matrix, start, midCol+1, cols-1, target)
	}
	if target <= matrix[start+1][midCol-1] {
		return binarySearch(matrix, start+1, 0, midCol-1, target)
	}
	return binarySearch(matrix, start+1, midCol+1, cols-1, target)
}
