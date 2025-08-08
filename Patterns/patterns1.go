package main

import (
	"fmt"
)

func main() {
	pattern(4)
}

func pattern(n int) {
	for row := 0; row <= n; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
