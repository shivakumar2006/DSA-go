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
		start = int(math.Max(float64(start), float64(arr[i])))
		end += arr[i]
	}

	for start < end {
		sum := 0
		pieces := 1
		mid := start + (end-start)/2
		for i := 0; i < len(arr); i++ {
			if sum+arr[i] > mid {
				sum = arr[i]
				pieces++
			} else {
				sum += arr[i]
			}
		}
		if pieces > m {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	arr := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(split(arr, m))
// }

// func split(arr []int, m int) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(arr); i++ {
// 		start = int(math.Max(float64(start), float64(arr[i])))
// 		end += arr[i]
// 	}

// 	for start < end {
// 		sum := 0
// 		pieces := 1
// 		mid := start + (end-start)/2
// 		for i := 0; i < len(arr); i++ {
// 			if sum+arr[i] > mid {
// 				sum = arr[i]
// 				pieces++
// 			} else {
// 				sum += arr[i]
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

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	arr := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(split(arr, m))
// }

// func split(arr []int, m int) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(arr); i++ {

// 		// find max number in the enite array
// 		start = int(math.Max(float64(start), float64(arr[i])))
// 		end += arr[i]
// 	}

// 	for start < end {
// 		sum := 0
// 		pieces := 1
// 		mid := start + (end-start)/2
// 		for i := 0; i < len(arr); i++ {
// 			if sum+arr[i] > mid {
// 				sum = arr[i]
// 				pieces++
// 			} else {
// 				sum += arr[i]
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

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	arr := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(split(arr, m))
// }

// func split(arr []int, m int) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(arr); i++ {
// 		start = int(math.Max(float64(start), float64(arr[i])))
// 		end += arr[i]
// 	}

// 	for start < end {
// 		sum := 0
// 		pieces := 1
// 		mid := start + (end-start)/2
// 		for i := 0; i < len(arr); i++ {
// 			if sum+arr[i] > mid {
// 				sum = arr[i]
// 				pieces++
// 			} else {
// 				sum += arr[i]
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

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	arr := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(split(arr, m))
// }

// func split(arr []int, m int) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(arr); i++ {
// 		start = int(math.Max(float64(start), float64(arr[i])))
// 		end += arr[i]
// 	}

// 	for start < end {
// 		sum := 0
// 		pieces := 1
// 		mid := start + (end-start)/2
// 		for i := 0; i < len(arr); i++ {
// 			if sum+arr[i] > mid {
// 				sum = arr[i]
// 				pieces++
// 			} else {
// 				sum += arr[i]
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

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	arr := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(split(arr, m))
// }

// func split(arr []int, m int) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(arr); i++ {
// 		start = int(math.Max(float64(start), float64(arr[i]))) // max element
// 		end += arr[i]                                          // sum of all element = 32
// 	}

// 	for start < end {
// 		mid := start + (end-start)/2 // 10 + (32 - 10)/2 -> 10+(22)/2 -> 10+11 -> 21
// 		sum := 0
// 		pieces := 1
// 		for i := 0; i < len(arr); i++ {
// 			if sum+arr[i] > mid {
// 				sum = arr[i]
// 				pieces++
// 			} else {
// 				sum += arr[i]
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

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	arr := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(split(arr, m)) // Output should be 18
// }

// func split(arr []int, m int) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(arr); i++ {
// 		start = int(math.Max(float64(start), float64(arr[i]))) // max element in array, minimum possible largest sum
// 		end += arr[i]                                          // sum of all elements, maximum possible largest sum
// 	}

// 	for start < end {
// 		mid := start + (end-start)/2 // middle value between start and end
// 		sum := 0
// 		pieces := 1 // number of subarrays

// 		for i := 0; i < len(arr); i++ {
// 			if sum+arr[i] > mid {
// 				// start new subarray
// 				sum = arr[i]
// 				pieces++
// 			} else {
// 				sum += arr[i]
// 			}
// 		}

// 		if pieces > m {
// 			// too many subarrays → increase allowed sum
// 			start = mid + 1
// 		} else {
// 			// valid → try to minimize sum
// 			end = mid
// 		}
// 	}

// 	return start // ✅ this is the minimum largest subarray sum
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	arr := []int{7, 2, 5, 10, 8}
// 	m := 2
// 	fmt.Println(split(arr, m))
// }

// func split(arr []int, m int) int {
// 	start, end := 0, 0
// 	for i := 0; i < len(arr); i++ {
// 		start = int(math.Max(float64(start), float64(arr[i]))) // find the largest number in the whole array...
// 		end += arr[i]                                          // 7 + 2 + 5 + 10 + 8 = 32
// 	}

// 	for start < end {
// 		mid := start + (end-start)/2 // (start + end) / 2 -> (10+32)/2 -> 42/2 -> 21
// 		sum := 0
// 		pieces := 1 // subarrays
// 		for i := 0; i < len(arr); i++ {
// 			if sum+arr[i] > mid { // 7 = 7, 7+2=9, 9+5=14, 14+10=24(invalid)
// 				sum = arr[i]
// 				pieces++ // now we have 3 subarrays
// 			} else {
// 				sum += arr[i]
// 			}
// 		}
// 		if pieces > m { // if the subarrays > 2 then increase mid and go to next iteration
// 			start = mid + 1
// 		} else {
// 			end = mid // otherwise return mid
// 		}
// 	}
// 	return -1
// }

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
