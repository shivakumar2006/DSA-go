// products of the digits...

package main

import "fmt"

func main() {
	fmt.Println(prod(1342))
}

func prod(n int) int {
	if n%10 == n {
		return n
	}

	return (n % 10) * prod(n/10)
}

// package main

// import "fmt"

// func main() {
// 	ans := prod(1342)
// 	fmt.Println(ans)
// }

// func prod(n int) int {
// 	if n%10 == n {
// 		return n
// 	}

// 	return (n % 10) * prod(n/10)
// }

// package main

// import "fmt"

// func main() {
// 	ans := prod(1342)
// 	fmt.Println(ans)
// }

// func prod(n int) int {
// 	if n%10 == n {
// 		return n
// 	}

// 	return (n % 10) * prod(n/10)
// }
