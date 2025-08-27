package main

import "fmt"

func main() {
	n := 4
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}

	knight(board, 0, 0, 4)
}

func knight(board [][]bool, row, col int, knights int) {
	if knights == 0 {
		display(board)
		fmt.Println()
		return 
	}

	if row == len(board) {
		return 
	}

	if col == len(board) {
		knight(board, row+1, 0, knights)
		return 
	}

	if isSafe(board, row, col) {
		board[row][col] = true 
		knight(board, row, col+1, knights-1) 
		board[row][col] = false 
	}
	knight(board, row, col+1, knights)
}

func isSafe(board [][]bool, row, col int) bool {
	if isValid(board, row-2, col-1) && board[row-2][col-1] {
		return false 
	}
	if isValid(board, row-1, col-1, )

	return true 
}

func display(board [][]bool) {
	for _, row := range board {
		for _, element := range row {
			if element {
				fmt.Print("K ")
			} else {
				fmt.Print("X ")
			}
		}
		fmt.Println()
	}
}

// func knight(board [][]bool, row, col int, knights int) {
// 	if knights == 0 {
// 		display(board)
// 		fmt.Println()
// 		return
// 	}

// 	// if our colun is out of bound
// 	if row == len(board) {
// 		return
// 	}

// 	if col == len(board) {
// 		knight(board, row+1, 0, knights)
// 		return
// 	}

// 	if isSafe(board, row, col) {
// 		board[row][col] = true
// 		knight(board, row, col+1, knights-1)
// 		board[row][col] = false
// 	}

// 	knight(board, row, col+1, knights)
// }

// func isSafe(board [][]bool, row, col int) bool {
// 	if isValid(board, row-2, col-1) && board[row-2][col-1] {
// 		return false
// 	}
// 	if isValid(board, row-2, col+1) && board[row-2][col+1] {
// 		return false
// 	}
// 	if isValid(board, row-1, col-2) && board[row-1][col-2] {
// 		return false
// 	}
// 	if isValid(board, row-1, col+2) && board[row-1][col+2] {
// 		return false
// 	}
// 	return true
// }

// func isValid(board [][]bool, row, col int) bool {
// 	if row >= 0 && row < len(board) && col >= 0 && col < len(board) {
// 		return true
// 	}

// 	return false
// }

// func display(board [][]bool) {
// 	for _, row := range board {
// 		for _, element := range row {
// 			if element {
// 				fmt.Print("K ")
// 			} else {
// 				fmt.Print("X ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
