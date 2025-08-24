package main

import "fmt"

func main() {
	path("", 3, 3)
}

func path(process string, row, col int) {
	if row == 1 && col == 1 {
		fmt.Println(process)
		return
	}

	if row > 1 {
		path(process+"D", row-1, col)
	}

	if col > 1 {
		path(process+"R", row, col-1)
	}
}

// func path(process string, row, col int) {
// 	if row == 1 && col == 1 {
// 		fmt.Println(process)
// 		return
// 	}

// 	if row > 1 {
// 		path(process+"D", row-1, col)
// 	}

// 	if col > 1 {
// 		path(process+"R", row, col-1)
// 	}
// }

// func path(process string, row, col int) {
// 	if row == 1 && col == 1 {
// 		fmt.Println(process)
// 		return
// 	}

// 	if row > 1 {
// 		path(process+"D", row-1, col) // go row by row
// 	}

// 	if col > 1 {
// 		path(process+"R", row, col-1) // go col by col
// 	}
// }
