// Maze woth obstacle

package main

import "fmt"

func main() {
	maze := [][]bool{
		{true, true, true},
		{true, false, true},
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

	if row < len(maze)-1 {
		path(process+"D", maze, row+1, col)
	}

	if col < len(maze[0])-1 {
		path(process+"R", maze, row, col+1)
	}
}

// func path(process string, maze [][]bool, row, col int) {
// 	if row == len(maze)-1 && col == len(maze[0])-1 {
// 		fmt.Println(process)
// 		return
// 	}

// 	if !maze[row][col] {
// 		return
// 	}

// 	if row < len(maze)-1 {
// 		path(process+"D", maze, row+1, col)
// 	}
// 	if col < len(maze[0])-1 {
// 		path(process+"R", maze, row, col+1)
// 	}
// }

// func path(process string, maze [][]bool, row, col int) {
// 	if row == len(maze)-1 && col == len(maze[0])-1 {
// 		fmt.Println(process)
// 		return
// 	}

// 	if !maze[row][col] {
// 		return
// 	}

// 	if row < len(maze)-1 {
// 		path(process+"D", maze, row+1, col)
// 	}
// 	if col < len(maze[0])-1 {
// 		path(process+"R", maze, row, col+1)
// 	}
// }
