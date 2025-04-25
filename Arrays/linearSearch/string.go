package main

import (
	"fmt"
)

func main() {
	str := "shiva"
	target := "v"

	res := findChar(str, target)
	if res != -1 {
		fmt.Printf("Element is found at the index of %d\n", res)
	} else {
		fmt.Println("Element is not found")
	}
}

func findChar(str string, target string) int {
	for i := 0; i <= len(str); i++ {
		if string(str[i]) == target {
			return i
		}
	}
	return -1
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	str := "shiva"
// 	targetChar := "v"

// 	resultStr := search(str, targetChar)
// 	if resultStr != -1 {
// 		fmt.Printf("Element is found at the index of %d\n", resultStr)
// 	} else {
// 		fmt.Println("Element is not found")
// 	}
// }

// func search(str string, target string) int {
// 	for i := 0; i <= len(str); i++ {
// 		if string(str[i]) == target {
// 			return i
// 		}
// 	}
// 	return -1
// }
