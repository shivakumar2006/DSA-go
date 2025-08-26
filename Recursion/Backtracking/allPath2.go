package main

import "fmt"

func main() {
	maze := [][]bool{
		{true, true, true},
		{true, true, true},
		{true, true, true},
	}

	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}

	allPath("", maze, 0, 0, steps, 1)
}

func allPath(process string, maze [][]bool, row, col int, steps [][]int, step int) {
	if row == len(maze)-1 && col == len(maze[0])-1 {
		steps[row][col] = step
		for _, arr := range steps {
			fmt.Println(arr)
		}
		fmt.Println(process)
		fmt.Println()
		return
	}

	if !maze[row][col] {
		return
	}

	maze[row][col] = false
	steps[row][col] = step

	if row < len(maze)-1 {
		allPath(process+"D", maze, row+1, col, steps, step+1)
	}
	if col < len(maze[0])-1 {
		allPath(process+"R", maze, row, col+1, steps, step+1)
	}
	if row > 0 {
		allPath(process+"U", maze, row-1, col, steps, step+1)
	}
	if col > 0 {
		allPath(process+"L", maze, row, col-1, steps, step+1)
	}

	maze[row][col] = true
	steps[row][col] = step
}

// package main

// import "fmt"

// func main() {
// 	maze := [][]bool{
// 		{true, true, true},
// 		{true, true, true},
// 		{true, true, true},
// 	}

// 	// create steps board
// 	steps := make([][]int, len(maze))
// 	for i := range steps {
// 		steps[i] = make([]int, len(maze[0]))
// 	}

// 	path("", maze, 0, 0, steps, 1)
// }

// func path(process string, maze [][]bool, row, col int, steps [][]int, step int) {
// 	if row == len(maze)-1 && col == len(maze[0])-1 {
// 		steps[row][col] = step
// 		// print the steps board
// 		for _, arr := range steps {
// 			fmt.Println(arr)
// 		}
// 		fmt.Println(process)
// 		fmt.Println()
// 		return
// 	}

// 	if !maze[row][col] {
// 		return
// 	}

// 	maze[row][col] = false
// 	steps[row][col] = step

// 	if row < len(maze)-1 {
// 		path(process+"D", maze, row+1, col, steps, step+1)
// 	}
// 	if col < len(maze[0])-1 {
// 		path(process+"R", maze, row, col+1, steps, step+1)
// 	}
// 	if row > 0 {
// 		path(process+"U", maze, row-1, col, steps, step+1)
// 	}
// 	if col > 0 {
// 		path(process+"L", maze, row, col-1, steps, step+1)
// 	}

// 	maze[row][col] = true
// 	steps[row][col] = 0
// }
