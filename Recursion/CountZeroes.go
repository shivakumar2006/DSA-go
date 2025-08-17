package main

import "fmt"

func main() {
	fmt.Println(count(3100002))
}

func count(n int) int {
	return helper(n, 0)
}

func helper(n int, count int) int {
	if n == 0 {
		return count
	}

	rem := n % 10
	if rem == 0 {
		return helper(n/10, count+1)
	}
	return helper(n/10, count)
}

// package main

// import "fmt"

// func main() {
// 	fmt.Println(count(310010))
// }

// func count(n int) int {
// 	return helper(n, 0)
// }

// func helper(n int, count int) int {
// 	if n == 0 {
// 		return count
// 	}

// 	rem := n % 10
// 	if rem == 0 {
// 		return helper(n/10, count+1)
// 	}
// 	return helper(n/10, count)
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println(count(310010))
// }

// func count(n int) int {
// 	return helper(n, 0)
// }

// func helper(n int, count int) int {
// 	if n == 0 {
// 		return count
// 	}

// 	rem := n % 10
// 	if rem == 0 {
// 		return helper(n/10, count+1)
// 	}
// 	return helper(n/10, count)
// }

// func count(n int) int {
// 	return helper(n, 0)
// }

// func helper(n int, c int) int {
// 	if n == 0 {
// 		return c
// 	}

// 	rem := n % 10
// 	if rem == 0 {
// 		return helper(n/10, c+1)
// 	}
// 	return helper(n/10, c)
// }

// package main

// import "fmt"

// func main() {
// 	fmt.Println(count(310210))
// }

// func count(n int) int {
// 	return helper(n, 0)
// }

// func helper(n int, c int) int {
// 	if n == 0 {
// 		return c
// 	}

// 	rem := n % 10
// 	if rem == 0 {
// 		return helper(n/10, c+1)
// 	}
// 	return helper(n/10, c)
// }
