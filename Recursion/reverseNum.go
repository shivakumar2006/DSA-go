// Reverse a number

// way-2
// this approach is bit slower
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rev(1234))
}

func rev(n int) int {
	if n == 0 {
		return 0
	}

	digits := int(math.Log10(float64(n))) // number of digits - 1
	return (n%10)*int(math.Pow(10, float64(digits))) + rev(n/10)
}

// way-1
// package main

// import "fmt"

// var sum int = 0

// func main() {
// 	rev(1342)
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
