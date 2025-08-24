// solve that same question using slice

package main

import "fmt"

func main() {
	ans := path("", 3, 3)
	fmt.Println(ans)
}

func path(process string, row, col int) []string {
	if row == 1 && col == 1 {
		return []string{process}
	}

	result := []string{}
	if row > 1 {
		down := path(process+"D", row-1, col)
		result = append(result, down...)
	}

	if col > 1 {
		right := path(process+"R", row, col-1)
		result = append(result, right...)
	}
	return result
}
