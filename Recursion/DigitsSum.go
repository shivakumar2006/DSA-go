// Sum of all the digits

package main

import "fmt"

func main() {
	fmt.Println(sum(1422))
}

func sum(n int) int {
	if n == 0 {
		return 0
	}

	return (n % 10) + sum(n/10)
}

// package main

// import "fmt"

// func main() {
// 	ans := sum(1422)
// 	fmt.Println(ans)
// }

// func sum(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	return (n % 10) + sum(n/10)
// }

// package main

// import "fmt"

// func main() {
// 	ans := sum(1342)
// 	fmt.Println(ans)
// }

// func sum(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	return (n % 10) + sum(n/10)
// }
