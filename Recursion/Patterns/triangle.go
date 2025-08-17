package main

import "fmt"

func main() {
	triangle(5, 0)
}

func triangle(row int, col int) {
	if row == 0 {
		return
	}

	if col < row {
		triangle(row, col+1)
		fmt.Print("* ")
	} else {
		triangle(row-1, 0)
		fmt.Println()
	}
}

// func triangle(row int, col int) {
// 	if row == 0 {
// 		return
// 	}

// 	if col < row {
// 		fmt.Print("* ")
// 		triangle(row, col+1)
// 	} else {
// 		fmt.Println()
// 		triangle(row-1, 0)
// 	}
// }
