// find the factorial of all the numbers...

package main

import "fmt"

func main() {
	ans := fact(5)
	fmt.Println(ans)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

// package main

// import "fmt"

// func main() {
// 	ans := fact(5)
// 	fmt.Println("factorial number : ", ans)
// }

// func fact(n int) int {
// 	if n < 1 {
// 		return 0
// 	}

// 	return n + fact(n-1)
// }

// package main

// import "fmt"

// func main() {
// 	ans := fact(5)
// 	fmt.Println("factorial of number : ", ans)
// }

// func fact(n int) int {
// 	if n < 1 {
// 		return 1
// 	}

// 	return n + fact(n-1)
// }
