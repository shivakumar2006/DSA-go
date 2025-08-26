package main

import "fmt"

func main() {
	maze := [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
	}
	path("", maze, 0, 0)
}

func path(process string, maze [][]bool, row, col int) {
	if row == len(maze)-1 && col == len(maze[0])-1 {
		fmt.Println(process)
		return
	}

	if !maze[row][col] {
		return
	}

	maze[row][col] = false

	if row < len(maze)-1 {
		path(process+"D", maze, row+1, col)
	}
	if col < len(maze[0])-1 {
		path(process+"R", maze, row, col+1)
	}
	if row > 0 {
		path(process+"U", maze, row-1, col)
	}
	if col > 0 {
		path(process+"L", maze, row, col-1)
	}

	maze[row][col] = true
}

// func path(process string, maze [][]bool, row, col int) {
// 	if row == len(maze)-1 && col == len(maze[0])-1 {
// 		fmt.Println(process)
// 		return
// 	}

// 	if !maze[row][col] {
// 		return
// 	}

// 	// i am considering this block in my path
// 	maze[row][col] = false

// 	if row < len(maze)-1 {
// 		path(process+"D", maze, row+1, col)
// 	}

// 	if col < len(maze[0])-1 {
// 		path(process+"R", maze, row, col+1)
// 	}

// 	if row > 0 {
// 		path(process+"U", maze, row-1, col)
// 	}

// 	if col > 0 {
// 		path(process+"L", maze, row, col-1)
// 	}

// 	// Backtrack unmark this cell
// 	maze[row][col] = true
// }
