package main

import "fmt"

func main() {
	maze := count(3, 3)
	fmt.Println(maze)
}

func count(row, col int) int {
	if row == 1 || col == 1 {
		return 1
	}

	left := count(row-1, col)  // here when we go on down means row by row
	right := count(row, col-1) // here when we go right means col by col
	return left + right
}

// the final output is 6 that means we have 6 ways to go to the 1 place to the tar place...
