// Reverse a number

package main

import "fmt"

var sum int = 0

func main() {
	rev(1342)
	fmt.Println(sum)
}

func rev(n int) {
	if n == 0 {
		return
	}

	remainder := n % 10
	sum = sum*10 + remainder
	rev(n / 10)
}

// WAY - 1 ...
// package main

// import "fmt"

// var sum int = 0

// func main() {
// 	rev(1824)
// 	fmt.Println(sum)
// }

// func rev(n int) {
// 	if n == 0 {
// 		return
// 	}
// 	remainder := n % 10
// 	sum = sum*10 + remainder
// 	rev(n / 10)
// }
