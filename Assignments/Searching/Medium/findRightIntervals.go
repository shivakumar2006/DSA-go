// Find Right Interval
// https://leetcode.com/problems/find-right-interval/description/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{3, 4}, {2, 3}, {1, 2}}
	fmt.Println(findRightInterval(arr))
}

func findRightInterval(arr [][]int) []int {
	n := len(arr)
	if n == 1 {
		if arr[0][0] >= arr[0][1] {
			return []int{0}
		}
		return []int{-1}
	}

	start := make([]int, n)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		start[i] = arr[i][0]
		indices[i] = i
	}

	// sort start
	sort.Slice(indices, func(i, j int) bool {
		return start[indices[i]] < start[indices[j]]
	})

	result := make([]int, n)
	for i := 0; i < n; i++ {
		end := arr[i][1]

		l, r := 0, n-1
		ans := -1
		for l <= r {
			mid := l + (r-l)/2
			if start[indices[mid]] >= end {
				ans = indices[mid]
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		result[i] = ans
	}
	return result
}
