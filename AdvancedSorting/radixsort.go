// Radix Sort

package main

import "fmt"

func RadixSort(arr []int) {
	max := maxValue(arr)

	for exp := 1; max/exp > 0; exp *= 10 {
		CountSort(arr, exp)
	}
}

func CountSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)

	// count
	for i := 0; i < n; i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	fmt.Println("\nCount array for", exp, ";", count)

	// update count aray
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	fmt.Print("Updated count array : ", count)

	// build output array
	for i := n - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	fmt.Println("Output array : ", output)

	copy(arr, output)
}

func maxValue(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	arr := []int{29, 83, 471, 36, 91, 8}
	fmt.Println("Original Array : ", arr)
	RadixSort(arr)
	fmt.Println("orted array : ", arr)
}
