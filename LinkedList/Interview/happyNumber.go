// Happy Number
// Google question

package main

import "fmt"

func isHappy(n int) bool {
	fast := n
	slow := n

	for {
		fast = findSquare(findSquare(fast))
		slow = findSquare(slow)

		if slow == 1 && fast == 1 {
			return true
		}

		if slow == fast {
			break
		}
	}

	return false

}

func findSquare(num int) int {
	ans := 0

	for num > 0 {
		rem := num % 10
		ans += rem * rem
		num = num / 10
	}

	return ans
}

func main() {
	n := 12
	fmt.Println(isHappy(n))
}
