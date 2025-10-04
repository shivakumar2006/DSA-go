// Median of 2 Sorted array
// https://leetcode.com/problems/median-of-two-sorted-arrays/description/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	// ensure nums1 is the smallest array
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m
	half := (m + n + 1) / 2 // 2 + 1 + 1 = (4 / 2) = 2

	// safe sentinel that is outside the problem constraints (-1e6..1e6)
	const INF = int(1e9)

	for low <= high {
		i := (low + high) / 2
		j := half - i

		maxLeft1 := -INF
		if i > 0 {
			maxLeft1 = nums1[i-1]
		}
		minRight1 := INF
		if i < m {
			minRight1 = nums1[i]
		}

		maxLeft2 := -INF
		if j > 0 {
			maxLeft2 = nums2[j-1]
		}
		minRight2 := INF
		if j < n {
			minRight2 = nums2[j]
		}

		// found correct partition
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 0 {
				leftMax := max(maxLeft1, maxLeft2)
				rightMin := min(minRight1, minRight2)
				return float64(leftMax+rightMin) / 2.0
			}
			return float64(max(maxLeft1, maxLeft2))
		}

		// move partition
		if maxLeft1 > minRight2 {
			high = i - 1
		} else {
			low = i + 1
		}
	}

	return 0.0
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
