package main

import (
	"fmt"
)

func main() {
	arr := []int{18, 124, 9, 1764, 98, 1, 69}
	res := oddNum(arr)
	fmt.Printf("Count of the number of digit is : %d\n", res)
}

func oddNum(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if oddDigits(arr[i]) {
			count++
		}
	}
	return count
}

func oddDigits(num int) bool {
	if num < 0 {
		num = -num
	}

	digits := 0
	for num != 0 {
		num = num / 10
		digits++
	}
	return digits%2 != 0
}

// func oddNum(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if oddDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func oddDigits(nums int) bool {
// 	if nums < 0 {
// 		nums = -nums
// 	}

// 	digits := 0
// 	for nums != 0 {
// 		nums = nums / 10
// 		digits++
// 	}
// 	return digits%2 != 0
// }

// func oddNum(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if oddDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func oddDigits(nums int) bool {
// 	if nums < 0 {
// 		nums = -nums
// 	}

// 	digits := 0
// 	for nums != 0 {
// 		nums = nums / 10
// 		digits++
// 	}
// 	return digits%2 != 0
// }

// func oddNum(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if oddDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func oddDigits(num int) bool {
// 	if num < 0 {
// 		num = -num
// 	}

// 	digits := 0
// 	for num != 0 {
// 		num = num / 10
// 		digits++
// 	}
// 	return digits%2 != 0
// }

// func oddNum(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if oddDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func oddDigits(num int) bool {
// 	if num < 0 {
// 		num = -num
// 	}

// 	digits := 0
// 	for num != 0 {
// 		num = num / 10
// 		digits++
// 	}
// 	return digits%2 != 0
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{18, 124, 9, 1764, 98, 1}

// 	res := oddNum(arr)
// 	fmt.Printf("Coount of the number of digits is : %d\n", res)
// }

// func oddNum(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if oddDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func oddDigits(num int) bool {
// 	if num < 0 {
// 		num = -num
// 	}

// 	digits := 0
// 	for num != 0 {
// 		num = num / 10
// 		digits++
// 	}
// 	return digits%2 != 0
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{18, 124, 9, 1764, 98, 1}

// 	res := oddNum(arr)
// 	fmt.Printf("Count of the number of digits : %d\n", res)
// }

// func oddNum(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if oddDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func oddDigits(num int) bool {
// 	if num < 0 {
// 		num = -num
// 	}

// 	digits := 0
// 	for num != 0 {
// 		num = num / 10
// 		digits++
// 	}
// 	return digits%2 != 0
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{18, 124, 9, 1764, 98, 1}

// 	res := oddNum(arr)
// 	fmt.Printf("Count of the numbers of digits : %d\n", res)
// }

// func oddNum(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if oddDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func oddDigits(num int) bool {
// 	if num < 0 {
// 		num = -num
// 	}

// 	digits := 0
// 	for num != 0 {
// 		num = num / 10
// 		digits++
// 	}
// 	return digits%2 != 0
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{18, 124, 9, 1764, 98, 1}

// 	res := oddNum(arr)
// 	fmt.Printf("Count of the number of odd digits : %d\n", res)
// }

// func oddNum(arr []int) int {
// 	count := 0
// 	for i := 0; i < len(arr); i++ {
// 		if oddDigits(arr[i]) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func oddDigits(num int) bool {
// 	if num < 0 {
// 		num = -num
// 	}

// 	digits := 0
// 	for num != 0 {
// 		num = num / 10
// 		digits++
// 	}
// 	return digits%2 != 0
// }
