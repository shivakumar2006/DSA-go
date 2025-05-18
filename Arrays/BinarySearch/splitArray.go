// Split array largest sum ...

package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{7, 2, 5, 10, 8}
	m := 2
	fmt.Println(split(arr, m))
}

func split(arr []int, m int) int {
	start, end := 0, 0
	for i := 0; i < len(arr); i++ {
		start = int(math.Max(float64(start), float64(arr[i]))) // find the largest number in the whole array...
		end += arr[i]                                          // 7 + 2 + 5 + 10 + 8 = 32
	}

	for start < end {
		mid := start + (end-start)/2 // (start + end) / 2 -> (10+32)/2 -> 42/2 -> 21
		sum := 0
		pieces := 1 // subarrays
		for i := 0; i < len(arr); i++ {
			if sum+arr[i] > mid { // 7 = 7, 7+2=9, 9+5=14, 14+10=24(invalid)
				sum = arr[i]
				pieces++ // now we have 3 subarrays
			} else {
				sum += arr[i]
			}
		}
		if pieces > m { // if the subarrays > 2 then increase mid and go to next iteration
			start = mid + 1
		} else {
			end = mid // otherwise return mid
		}
	}
	return -1
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	nums := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(split(nums, m)) // Output should be 18
// }

// func split(nums []int, m int) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(nums); i++ {
// 		start = int(math.Max(float64(start), float64(nums[i])))
// 		end += nums[i]
// 	}

// 	for start < end {
// 		mid := start + (end-start)/2

// 		sum := 0
// 		pieces := 1   // Pieces means subarrays

// 		for i := 0; i < len(nums); i++ {
// 			if sum+nums[i] > mid {
// 				sum = nums[i]
// 				pieces++
// 			} else {
// 				sum += nums[i]
// 			}
// 		}

// 		if pieces > m {
// 			start = mid + 1
// 		} else {
// 			end = mid
// 		}
// 	}

// 	return start
// }
