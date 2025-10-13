// Merge Intervals
// https://leetcode.com/problems/merge-intervals/description/

package main

import (
	"fmt"
	"sort"
)

func merge(arr [][]int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}

	// sort the start line
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	// start merging
	merged := [][]int{arr[0]}

	for i := 0; i < len(arr); i++ {
		last := merged[len(merged)-1]
		current := arr[i]

		// if overlapping
		if current[0] <= last[1] {
			// merge by updating an time
			if current[1] > last[1] {
				last[1] = current[1]
			}
			merged[len(merged)-1] = last
		} else {
			// no overlap -> start a new interval
			merged = append(merged, current)
		}
	}
	return merged
}

func main() {
	interval := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(interval))
}
