// Minimum Absolute Sum Difference
// https://leetcode.com/problems/minimum-absolute-sum-difference/description/

package main

import (
	"fmt"
	"sort"
)

func minAbsoluteSumDiff(nums1, nums2 []int) int {
	mod := int(1e9 + 7)
	n := len(nums1)

	// copy and sort nums1 for binarySearch
	sortedNums1 := make([]int, n)
	copy(sortedNums1, nums1)
	sort.Ints(sortedNums1)

	sumDiff := 0
	maxImprovement := 0

	for i := 0; i < n; i++ {
		diff := abs(nums1[i] - nums2[i])
		sumDiff = (sumDiff + diff) % mod

		// binarySearch for closest to nums2[i]
		index := sort.SearchInts(sortedNums1, nums2[i])

		// check candidate at index (>= nums2[i])
		if index < n {
			newDiff := abs(sortedNums1[index] - nums2[i]) // ✅ correct
			improvement := diff - newDiff
			if improvement > maxImprovement {
				maxImprovement = improvement
			}
		}
		// check candidate at index-1 (< nums2[i])
		if index > 0 {
			newDiff := abs(sortedNums1[index-1] - nums2[i]) // ✅ correct
			improvement := diff - newDiff
			if improvement > maxImprovement {
				maxImprovement = improvement
			}
		}
	}
	return (sumDiff - maxImprovement + mod) % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	nums1 := []int{1, 7, 5}
	nums2 := []int{2, 3, 5}
	fmt.Println(minAbsoluteSumDiff(nums1, nums2))
}

// func minAbsoluteSumDiff(nums1, nums2 []int) int {
// 	mod := int(1e9 + 7)
// 	n := len(nums1)

// 	// copy ans sort nums1 for binarySearch
// 	sortedNums1 := make([]int, n)
// 	copy(sortedNums1, nums1)
// 	sort.Ints(sortedNums1)

// 	sumDiff := 0
// 	maxImprovement := 0

// 	for i := 0; i < n; i++ {
// 		diff := abs(nums1[i] - nums2[i])
// 		sumDiff = (sumDiff + diff) % mod

// 		// find the closest number in nums1 to nums2
// 		index := sort.SearchInts(sortedNums1, nums2[i])

// 		if index < n {
// 			// candidate on the right or equal
// 			newDiff := abs(sortedNums1[index] - nums2[i])
// 			improvement := diff - newDiff
// 			if improvement > maxImprovement {
// 				maxImprovement = improvement
// 			}
// 		}
// 		if index > 0 {
// 			// candidate to the left
// 			newDiff := abs(sortedNums1[index-1] - nums2[i])
// 			improvement := diff - newDiff
// 			if improvement > maxImprovement {
// 				maxImprovement = improvement
// 			}
// 		}
// 	}
// 	return (sumDiff - maxImprovement + mod) % mod
// }

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func main() {
// 	nums1 := []int{1, 7, 5}
// 	nums2 := []int{2, 3, 5}
// 	fmt.Println(minAbsoluteSumDiff(nums1, nums2))
// }
