// pattern 3

package main

import (
	"fmt"
)

func main() {
	pattern(4)
}

func pattern(n int) {
	for row := n; row >= 1; row-- {
		for col := 1; col <= row; col++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

// PATTERN 2

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	pattern2(4)
// }

// func pattern2(n int) {
// 	for i := 0; i <= n; i++ {
// 		for j := 1; j <= n; j++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

// PATTERN 1...

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	pattern(4)
// }

// func pattern(n int) {
// 	for row := 0; row <= n; row++ {
// 		for col := 1; col <= row; col++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }
