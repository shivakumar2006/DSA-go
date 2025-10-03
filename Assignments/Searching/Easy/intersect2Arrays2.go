// intersection of two arrays
// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersect(nums1, nums2))
}

func intersect(nums1, nums2 []int) []int {
	freq := make(map[int]int)
	for _, val := range nums1 {
		freq[val]++
	}

	res := []int{}
	for _, val := range nums2 {
		if freq[val] > 0 {
			res = append(res, val)
			freq[val]--
		}
	}
	return res
}

// func intersect(nums1, nums2 []int) []int {
// 	freq := make(map[int]int)
// 	for _, val := range nums1 {
// 		freq[val]++
// 	}

// 	res := []int{}
// 	for _, val := range nums2 {
// 		if freq[val] > 0 {
// 			res = append(res, val)
// 			freq[val]--
// 		}
// 	}

// 	return res
// }
