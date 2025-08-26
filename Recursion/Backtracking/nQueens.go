package main

import (
	"fmt"
	"math"
)

func main() {
	n := 5
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}
	fmt.Println("total solutions : ", queens(board, 0))
}

func queens(board [][]bool, row int) int {
	if row == len(board) {
		display(board)
		fmt.Println()
		return 1
	}

	count := 0

	for col := 0; col < len(board); col++ {
		if isSafe(board, row, col) {
			board[row][col] = true
			count += queens(board, row+1)
			board[row][col] = false
		}
	}

	return count
}

func isSafe(board [][]bool, row, col int) bool {
	// check vertically
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}

	// check up left
	maxLeft := int(math.Min(float64(row), float64(col)))
	for i := 1; i <= maxLeft; i++ {
		if board[row-i][col-i] {
			return false
		}
	}

	// check up right
	maxRight := int(math.Min(float64(row), float64(len(board)-col-1)))
	for i := 1; i <= maxRight; i++ {
		if board[row-i][col+i] {
			return false
		}
	}

	return true
}

func display(board [][]bool) {
	for _, row := range board {
		for _, element := range row {
			if element {
				fmt.Print("Q ")
			} else {
				fmt.Print("X ")
			}
		}
		fmt.Println()
	}
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	n := 4
// 	board := make([][]bool, n)
// 	for i := range board {
// 		board[i] = make([]bool, n)
// 	}
// 	fmt.Println("Total solutions : ", queens(board, 0))
// }

// func queens(board [][]bool, row int) int {
// 	if row == len(board) {
// 		display(board)
// 		fmt.Println()
// 		return 1
// 	}

// 	count := 0

// 	for col := 0; col < len(board); col++ {
// 		if isSafe(board, row, col) {
// 			board[row][col] = true
// 			count += queens(board, row+1)
// 			board[row][col] = false
// 		}
// 	}

// 	return count
// }

// func isSafe(board [][]bool, row, col int) bool {
// 	// check vertically
// 	for i := 0; i < row; i++ {
// 		if board[i][col] {
// 			return false
// 		}
// 	}

// 	// check up left
// 	maxLeft := int(math.Min(float64(row), float64(col)))
// 	for i := 1; i <= maxLeft; i++ {
// 		if board[row-i][col-i] {
// 			return false
// 		}
// 	}

// 	// check up right
// 	maxRight := int(math.Min(float64(row), float64(len(board)-col-1)))
// 	for i := 1; i <= maxRight; i++ {
// 		if board[row-i][col+i] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func display(board [][]bool) {
// 	for _, row := range board {
// 		for _, element := range row {
// 			if element {
// 				fmt.Println("Q ")
// 			} else {
// 				fmt.Println("X ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	n := 4
// 	board := make([][]bool, n)
// 	for i := range board {
// 		board[i] = make([]bool, n)
// 	}

// 	fmt.Println("Total Solutions:", queens(board, 0))
// }

// func queens(board [][]bool, row int) int {
// 	if row == len(board) {
// 		display(board)
// 		fmt.Println()
// 		return 1
// 	}

// 	count := 0

// 	for col := 0; col < len(board); col++ {
// 		if isSafe(board, row, col) {
// 			board[row][col] = true
// 			count += queens(board, row+1)
// 			board[row][col] = false
// 		}
// 	}

// 	return count
// }

// func isSafe(board [][]bool, row, col int) bool {
// 	// vertical
// 	for i := 0; i < row; i++ {
// 		if board[i][col] {
// 			return false
// 		}
// 	}

// 	// diagonal left
// 	maxLeft := int(math.Min(float64(row), float64(col)))
// 	for i := 1; i <= maxLeft; i++ {
// 		if board[row-i][col-i] {
// 			return false
// 		}
// 	}

// 	// diagonal right
// 	maxRight := int(math.Min(float64(row), float64(len(board)-col-1)))
// 	for i := 1; i <= maxRight; i++ {
// 		if board[row-i][col+i] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func display(board [][]bool) {
// 	for _, row := range board {
// 		for _, element := range row {
// 			if element {
// 				fmt.Print("Q ")
// 			} else {
// 				fmt.Print("X ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
