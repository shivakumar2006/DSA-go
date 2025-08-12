package main

import "fmt"

func main() {
	pyramid(5)
}

func pyramid(n int) {
	for row := 1; row <= 2*n-1; row++ {
		totalCol := row
		if row > n {
			totalCol = n*2 - row
		}

		noOfSpaces := n - totalCol
		for s := 0; s < noOfSpaces; s++ {
			fmt.Print("  ")
		}

		for col := totalCol; col >= 1; col-- {
			fmt.Print(col, " ")
		}
		for col := 2; col <= totalCol; col++ {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
}

// func pyramid(n int) {
// 	for row := 1; row <= 2*n-1; row++ {
// 		totalCol := row
// 		if row > n {
// 			totalCol = 2*n - row
// 		}

// 		noOfSpaces := n - totalCol
// 		for space := 0; space < noOfSpaces; space++ {
// 			fmt.Print("  ")
// 		}

// 		for col := totalCol; col >= 1; col-- {
// 			fmt.Print(col, " ")
// 		}
// 		for col := 2; col <= totalCol; col++ {
// 			fmt.Print(col, " ")
// 		}
// 		fmt.Println()
// 	}
// }

// func pyramid(n int) {
// 	for row := 1; row <= n; row++ {
// 		for space := 0; space < n-row; space++ {
// 			fmt.Print("  ")
// 		}

// 		for col := row; col >= 1; col-- {
// 			fmt.Print(col, " ")
// 		}
// 		for col := 2; col <= row; col++ {
// 			fmt.Print(col, " ")
// 		}
// 		fmt.Println()
// 	}
// }

// func pyramid(n int) {
// 	for row := 1; row <= n; row++ {

// 		for space := 0; space < n-row; space++ {
// 			fmt.Print("  ")
// 		}

// 		for col := row; col >= 1; col-- {
// 			fmt.Print(col, " ")
// 		}
// 		for col := 2; col <= row; col++ {
// 			fmt.Print(col, " ")
// 		}

// 		fmt.Println()
// 	}
// }

// func pyramid(n int) {
// 	for row := 1; row <= 2*n; row++ {
// 		totalCol := row
// 		if row > n {
// 			totalCol = 2*n - row
// 		}

// 		noOfSpaces := n - totalCol
// 		for s := 0; s < noOfSpaces; s++ {
// 			fmt.Print(" ")
// 		}
// 		for col := 1; col <= totalCol; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func pyramid(n int) {
// 	for row := n; row >= 1; row-- {
// 		for space := 1; space <= n-row; space++ {
// 			fmt.Print(" ")
// 		}

// 		for col := 1; col <= row; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func pyramid(n int) {
// 	for row := 0; row <= 2*n; row++ {
// 		totalCol := row
// 		if row > n {
// 			totalCol = 2*n - row
// 		}

// 		noOfSpaces := n - totalCol
// 		for s := 0; s < noOfSpaces; s++ {
// 			fmt.Print(" ")
// 		}

// 		for col := 0; col < totalCol; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func pyramid(n int) {
// 	for row := 1; row <= n; row++ {
// 		for space := 1; space <= n-row; space++ {
// 			fmt.Print(" ")
// 		}
// 		// totalCol := row

// 		// if row > n {
// 		// 	totalCol = 2*n - row
// 		// }

// 		for col := 1; col <= row; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}

// 	for i := n - 1; i >= 1; i-- {
// 		for space := 0; space < n-i; space++ {
// 			fmt.Print(" ")
// 		}

// 		for col := 1; col <= i; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func pyramid(n int) {
// 	for row := n; row >= 1; row-- {
// 		for space := 0; space < n-row; space++ {
// 			fmt.Print(" ")
// 		}

// 		for col := 1; col <= row; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// inverted pyramid
// func pyramid(n int) {
// 	for row := n; row >= 1; row-- {
// 		for space := 0; space < n-row; space++ {
// 			fmt.Print(" ")
// 		}

// 		for col := 1; col <= row; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// func pyramid(n int) {
// 	for row := 1; row <= n; row++ {
// 		// print space
// 		for space := 1; space <= n-row; space++ {
// 			fmt.Print(" ")
// 		}

// 		// print starts
// 		for col := 1; col <= row; col++ {
// 			fmt.Print("* ")
// 		}

// 		fmt.Println()
// 	}
// }
