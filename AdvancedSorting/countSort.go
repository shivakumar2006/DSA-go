// count sorting

package main

import "fmt"

// countsort for non-negative numbers
func CountSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// find max value
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	countArray := make([]int, max+1)

	// count frequency
	for _, num := range arr {
		countArray[num]++
	}

	index := 0
	for i, count := range countArray {
		for count > 0 {
			arr[index] = i
			index++
			count--
		}
	}
}

// handles negative numbers using a map
func CountSortHash(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// find min and max
	min, max := arr[0], arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}

	index := 0
	for i := min; i <= max; i++ {
		count := countMap[i]
		for j := 0; j < count; j++ {
			arr[index] = i
			index++
		}
	}
}

func main() {
	arr := []int{6, 3, 10, 9, 2, 4, 9, 7}
	CountSortHash(arr)
	fmt.Println(arr)
}
