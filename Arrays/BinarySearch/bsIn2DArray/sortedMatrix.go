// searching the element in sorted matrix in binarySearch

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
	tar := 8
	res := search(matrix, tar)
	fmt.Println(res)
}

func search(matrix [][]int, tar int) []int {
	row := len(matrix)
	col := len(matrix[0]) // in case if the col is emptu

	if row == 0 || len(matrix[0]) == 0 {
		return []int{-1, -1}
	}
	if row == 1 {
		return binarySearch(matrix, 0, 0, col-1, tar)
	}

	start := 0
	end := row - 1
	midCol := col / 2

	for start < end-1 {
		mid := start + (end-start)/2
		if matrix[mid][midCol] == tar {
			return []int{mid, midCol}
		} else if matrix[mid][midCol] < tar {
			start = mid
		} else {
			end = mid
		}
	}

	if matrix[start][midCol] == tar {
		return []int{start, midCol}
	}
	if matrix[start+1][midCol] == tar {
		return []int{start + 1, midCol}
	}
	if tar <= matrix[start][midCol-1] {
		return binarySearch(matrix, start, 0, midCol-1, tar)
	}
	if tar >= matrix[start][midCol+1] && tar <= matrix[start][col-1] {
		return binarySearch(matrix, start, midCol+1, col-1, tar)
	}
	if tar <= matrix[start+1][midCol-1] {
		return binarySearch(matrix, start+1, 0, midCol-1, tar)
	}
	return binarySearch(matrix, start+1, midCol-1, col-1, tar)
}

func binarySearch(matrix [][]int, row, start, end, tar int) []int {
	for start <= end {
		mid := start + (end-start)/2
		if matrix[row][mid] == tar {
			return []int{row, mid}
		} else if matrix[row][mid] < tar {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return []int{-1, -1}
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	matrix := [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}
// 	target := 9
// 	result := search(matrix, target)
// 	fmt.Println(result) // Output: [2 2]
// }

// func binarySearch(matrix [][]int, row, start, end, target int) []int {
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if matrix[row][mid] == target {
// 			return []int{row, mid}
// 		}
// 		if matrix[row][mid] < target {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return []int{-1, -1}
// }

// func search(matrix [][]int, target int) []int {
// 	rows := len(matrix)
// 	if rows == 0 || len(matrix[0]) == 0 {
// 		return []int{-1, -1}
// 	}
// 	cols := len(matrix[0])

// 	if rows == 1 {
// 		return binarySearch(matrix, 0, 0, cols-1, target)
// 	}

// 	start := 0
// 	end := rows - 1
// 	midCol := cols / 2

// 	for start < end-1 {
// 		mid := start + (end-start)/2
// 		if matrix[mid][midCol] == target {
// 			return []int{mid, midCol}
// 		}
// 		if matrix[mid][midCol] < target {
// 			start = mid
// 		} else {
// 			end = mid
// 		}
// 	}

// 	if matrix[start][midCol] == target {
// 		return []int{start, midCol}
// 	}
// 	if matrix[start+1][midCol] == target {
// 		return []int{start + 1, midCol}
// 	}

// 	if target <= matrix[start][midCol-1] {
// 		return binarySearch(matrix, start, 0, midCol-1, target)
// 	}
// 	if target >= matrix[start][midCol+1] && target <= matrix[start][cols-1] {
// 		return binarySearch(matrix, start, midCol+1, cols-1, target)
// 	}
// 	if target <= matrix[start+1][midCol-1] {
// 		return binarySearch(matrix, start+1, 0, midCol-1, target)
// 	}
// 	return binarySearch(matrix, start+1, midCol+1, cols-1, target)
// }
