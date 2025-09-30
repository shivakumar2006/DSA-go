package main

import "fmt"

func main() {
	num := 8
	ans := mySqrt(num)
	fmt.Println("squareroot of number is : ", ans)
}

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	var ans int

	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
