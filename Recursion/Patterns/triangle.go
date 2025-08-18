package main

import "fmt"

func main() {
	triangle(5, 0)
}

func triangle(r, c int) {
	if r == 0 {
		return
	}

	if c < r {
		fmt.Print("* ")
		triangle(r, c+1)
	} else {
		fmt.Println()
		triangle(r-1, 0)
	}
}

// func triangle(r, c int) {
// 	if r == 0 {
// 		return
// 	}

// 	if c < r {
// 		fmt.Print("* ")
// 		triangle(r, c+1)
// 	} else {
// 		fmt.Println()
// 		triangle(r-1, 0)
// 	}
// }

// func triangle(r int, c int) {
// 	if r == 0 {
// 		return
// 	}

// 	if c < r {
// 		triangle(r, c+1)
// 		fmt.Print("* ")
// 	} else {
// 		triangle(r-1, 0)
// 		fmt.Println()
// 	}
// }

// func triangle(row int, col int) {
// 	if row == 0 {
// 		return
// 	}

// 	if col < row {
// 		triangle(row, col+1)
// 		fmt.Print("* ")
// 	} else {
// 		triangle(row-1, 0)
// 		fmt.Println()
// 	}
// }

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
