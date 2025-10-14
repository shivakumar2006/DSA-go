// find max and min using recursion
// https://www.geeksforgeeks.org/dsa/recursive-programs-to-find-minimum-and-maximum-elements-of-array/

package main

import "fmt"

func maxAndMin(arr []int, index int) (int, int) {
	if index == len(arr)-1 {
		return arr[index], arr[index]
	}
	// recursive call for the rest of the array
	maxRest, minRest := maxAndMin(arr, index+1)

	if arr[index] > maxRest {
		maxRest = arr[index]
	}
	if arr[index] < minRest {
		minRest = arr[index]
	}
	return maxRest, minRest
}

func main() {
	arr := []int{1, 4, 3, -5, -4, 8, 6}
	maxVal, minVal := maxAndMin(arr, 0)
	fmt.Println("Maximum : ", maxVal)
	fmt.Println("Minimum : ", minVal)
}
