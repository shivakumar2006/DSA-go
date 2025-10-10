// Sort Array by Increasing Frequency
// https://leetcode.com/problems/sort-array-by-increasing-frequency/description/

package main

import (
	"fmt"
	"sort"
)

func frequencySort(arr []int) []int {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num]++
	}

	sort.Slice(arr, func(i, j int) bool {
		if freq[arr[i]] != freq[arr[j]] {
			return freq[arr[i]] < freq[arr[j]]
		}
		return arr[i] > arr[j]
	})
	return arr
}

func main() {
	arr := []int{1, 1, 2, 2, 2, 3}
	fmt.Println(frequencySort(arr))
}
