// iintersection of 2 arrays
// https://leetcode.com/problems/intersection-of-two-arrays/description/

package main

import "fmt"

func intersection(nums1, nums2 []int) []int {
	set := make(map[int]struct{})
	for _, val := range nums1 {
		set[val] = struct{}{}
	}

	resSet := make(map[int]struct{})
	for _, val := range nums2 {
		resSet[val] = struct{}{}
	}

	res := make([]int, 0, len(resSet))
	for v := range resSet {
		res = append(res, v)
	}

	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
