// Guess numner is higher or lower
package main

import "fmt"

func main() {
	n := 10
	fmt.Println("Picked number is : ", guessNumber(n))
}

var pick int = 6

func guess(num int) int {
	if num > pick {
		return -1
	} else if num < pick {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	start := 1
	end := n

	for start < end {
		mid := start + (end-start)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res < 0 {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

// func guess(num int) int {
// 	if num > pick {
// 		return -1
// 	} else if num < pick {
// 		return 1
// 	}
// 	return 0
// }

// func guessNumber(n int) int {
// 	start := 1
// 	end := n

// 	for start <= end {
// 		mid := start + (end-start)/2
// 		res := guess(mid)
// 		if res == 0 {
// 			return mid
// 		} else if res < 0 {
// 			end = mid - 1
// 		} else {
// 			start = mid + 1
// 		}
// 	}
// 	return -1
// }
