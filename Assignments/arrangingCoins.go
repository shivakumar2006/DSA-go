// Aranging coins
// https://leetcode.com/problems/arranging-coins/description/

package main

import "fmt"

func main() {
	n := 5
	fmt.Println("full rows : ", arrangeCoins(n))
}

func arrangeCoins(n int) int {
	start, end := 0, n
	for start <= end {
		mid := start + (end-start)/2
		k := mid*(mid+1)/2 <= n
		if k == true {
			start = mid + 1
		} else if k == false {
			end = mid - 1
		}
	}
	return end
}
