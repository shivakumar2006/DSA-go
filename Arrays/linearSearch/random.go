package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100, 500, 320}
	tar := 500
	ans := find(arr, tar)
	fmt.Println("the index of the targer element is : ", ans)
}

func find(arr []int, tar int) int {
	for i := 0; i <= len(arr); i++ {
		if arr[i] == tar {
			return i
		}
	}
	return -1
}

// this question is complete
