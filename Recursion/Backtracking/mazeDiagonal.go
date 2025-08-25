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
	if row > 1 && col > 1 {
		diagonal := path(process+"D", row-1, col-1)
		result = append(result, diagonal...)
	}

	if row > 1 {
		horizontal := path(process+"H", row-1, col)
		result = append(result, horizontal...)
	}

	if col > 1 {
		vertical := path(process+"V", row, col-1)
		result = append(result, vertical...)
	}
	return result
}

// func path(process string, row, col int) {
// 	if row == 1 && col == 1 {
// 		fmt.Println(process)
// 		return
// 	}

// 	if row > 1 && col > 1 {
// 		path(process+"D", row-1, col-1)
// 	}

// 	if row > 1 {
// 		path(process+"H", row-1, col)
// 	}

// 	if col > 1 {
// 		path(process+"V", row, col-1)
// 	}
// }
