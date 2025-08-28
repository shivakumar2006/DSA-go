package main

import "fmt"

const N = 9

func main() {
	board := [N][N]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if solveSudoku(&board) {
		display(board)
		fmt.Println("sudoku solve")
	} else {
		fmt.Println("not solve")
	}
}

func solveSudoku(board *[N][N]int) bool {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if board[row][col] == 0 {
				for num := 1; num <= 9; num++ {
					if isSafe(*board, row, col, num) {
						board[row][col] = num
						if solveSudoku(board) {
							return true
						}
						board[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isSafe(board [N][N]int, row, col, num int) bool {
	//check row
	for x := 0; x < N; x++ {
		if board[row][x] == num {
			return false
		}
	}

	// check col
	for x := 0; x < N; x++ {
		if board[x][col] == num {
			return false
		}
	}

	//check 3x3 subgrid
	startRow := row - row%3
	startCol := col - col%3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	// if the same number are not repeat on the row and col and in the 3x3 grid as well then return true
	return true
}

func display(board [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}
