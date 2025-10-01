// Maximum value at a given index in a Boundary array
// https://leetcode.com/problems/maximum-value-at-a-given-index-in-a-bounded-array/description/

package main

import "fmt"

func main() {
	n := 4
	index := 2
	maxSum := 6
	ans := maxValue(n, index, maxSum)
	fmt.Println(ans)
}

func maxValue(n, index, maxSum int) int {
	start, end := 0, maxSum
	ans := 1

	for start <= end {
		mid := start + (end-start)/2
		if isPossible(mid, n, index, maxSum) {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return ans
}

func isPossible(mid, n, index, maxSum int) bool {
	leftLen := index
	rightLen := n - index - 1

	left := calcSum(mid, leftLen)
	right := calcSum(mid, rightLen)

	total := left + right + maxSum
	return total <= maxSum
}

func calcSum(mid int, length int) int {
	if mid > length {
		last := mid - length
		return (mid - 1 + last) * length / 2
	} else {
		return (mid-1)*mid/2 + (length - (mid - 1))
	}
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func maxValue(n int, index int, maxSum int) int {
// 	maxVal := 1

// 	for val := 1; ; val++ {
// 		arr := make([]int, n)
// 		arr[index] = val

// 		// fill the left side
// 		for i := index - 1; i >= 0; i-- {
// 			arr[i] = max(1, arr[i+1]-1)
// 		}

// 		// fill right side
// 		for i := index + 1; i < n; i++ {
// 			arr[i] = max(1, arr[i-1]-1)
// 		}

// 		// calculate sum
// 		sum := 0
// 		for i := 0; i < n; i++ {
// 			sum += arr[i] // sum = sum + arr[i]
// 		}
// 		if sum > maxSum {
// 			break
// 		}
// 		maxVal = val
// 	}

// 	return maxVal
// }

// func maxValue(n int, index int, maxSum int) int {
// 	start, end := 1, maxSum
// 	ans := 1
// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if isPossible(mid, n, index, maxSum) {
// 			ans = mid
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return ans
// }

// func isPossible(mid int, n int, index int, maxSum int) bool {
// 	leftLen := index
// 	rightLen := n - index - 1

// 	leftSum := calcSum(mid, leftLen)
// 	rightSum := calcSum(mid, rightLen)

// 	total := leftSum + rightSum + mid
// 	return total <= maxSum
// }

// func calcSum(mid int, length int) int {
// 	if mid > length {
// 		last := mid - length
// 		return (mid - 1 + last) * length / 2
// 	} else {
// 		return (mid-1)*mid/2 + (length - (mid - 1))
// 	}
// }
