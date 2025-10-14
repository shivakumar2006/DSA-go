// sum traingle from array
// https://www.geeksforgeeks.org/dsa/sum-triangle-from-array/

package main

import "fmt"

func printTriangle(arr []int) {
	if len(arr) == 1 {
		fmt.Println(arr)
		return
	}

	temp := make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		x := arr[i] + arr[i+1]
		temp[i] = x
	}

	printTriangle(temp)

	fmt.Println(arr, " ")
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	printTriangle(arr)
}
