package main

import (
	"fmt"
	"math"
)

func main() {
	num := 1824
	if isPalindrome(num) {
		fmt.Println("palindrome : ", num)
	} else {
		fmt.Println("not a palindrome : ", num)
	}

	num2 := 1221
	if isPalindrome(num2) {
		fmt.Println("palindrome : ", num2)
	} else {
		fmt.Println("not a palindrome : ", num2)
	}
}

func rev(n int) int {
	if n == 0 {
		return 0
	}

	digits := int(math.Log10(float64(n)))
	return (n%10)*int(math.Pow(10, float64(digits))) + rev(n/10)
}

func isPalindrome(n int) bool {
	return n == rev(n)
}

// func main() {
// 	num := 1824
// 	if isPalindrome(num) {
// 		fmt.Println("palindrome : ", num)
// 	} else {
// 		fmt.Println("not a palindrome : ", num)
// 	}

// 	num2 := 1221
// 	if isPalindrome(num2) {
// 		fmt.Println("palindrome : ", num2)
// 	} else {
// 		fmt.Println("not a palindrome : ", num2)
// 	}
// }

// func rev(n int) int {
// 	if n == 0 {
// 		return 0
// 	}

// 	digits := int(math.Log10(float64(n)))
// 	return (n%10)*int(math.Pow(10, float64(digits))) + rev(n/10)
// }

// func isPalindrome(n int) bool {
// 	return n == rev(n)
// }
