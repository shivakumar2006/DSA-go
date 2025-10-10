// 3sum
// https://leetcode.com/problems/3sum/description/

package main

import (
	"fmt"
	"sort"
)

func threeSum(arr []int) [][]int {
	sort.Ints(arr)
	n := len(arr)
	res := [][]int{}

	for i := 0; i < n-2; i++ {
		if i > 0 && arr[i] == arr[i-1] { // skip duplicates
			continue
		}

		left, right := i+1, n-1

		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == 0 {
				res = append(res, []int{arr[i], arr[left], arr[right]})
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(arr))
}
