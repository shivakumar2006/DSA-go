// // find number of numbers that have even numbers of digits...

package main

import (
	"fmt"
)

func main() {
	arr := []int{18, 124, 9, 1764, 98, 1}

	res := evenNumbers(arr)
	fmt.Printf("Count of the numbers with even number of digits : %d\n", res)
}

func evenNumbers(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if evenDigits(arr[i]) {
			count++
		}
	}
	return count
}

func evenDigits(num int) bool {
	if num < 0 {
		num = -num
	}

	digits := 0
	for num != 0 {
		num = num / 10
		digits++
	}
	return digits%2 == 0
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{18, 124, 9, 1764, 98, 1}

// 	res := evenNumbers(arr)
// 	fmt.Printf("Count of the numbers with even numbers of digit : %d\n", res)
// }

// func evenNumbers(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if evenDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func evenDigits(num int) bool {
// 	if num < 0 {
// 		num = -num
// 	}

// 	digits := 0
// 	for num != 0 {
// 		num = num / 10 // just remove 1 digit on every iteration and add that 1 evertime until any number is not left for remoing
// 		digits++
// 	}
// 	return digits%2 == 0
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{18, 124, 9, 1764, 98, 1}

// 	result := evenNumbers(arr)
// 	fmt.Printf("count of numners with even numbers of digit : %d\n", result)
// }

// func evenNumbers(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if evenDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func evenDigits(num int) bool {
// 	if num < 0 {
// 		num = -num
// 	}

// 	digits := 0
// 	for num != 0 {
// 		num = num / 10
// 		digits++
// 	}
// 	return digits%2 == 0
// }
