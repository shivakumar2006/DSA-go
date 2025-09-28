// Intersection of 2 arrays
// https://leetcode.com/problems/intersection-of-two-arrays/description/

package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	for _, v := range nums1 {
		set[v] = struct{}{}
	}

	resSet := make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			resSet[v] = struct{}{}
		}
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

	fmt.Println("Intersection of 2 arrays is : ", intersection(nums1, nums2))
}
