package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5}
	tar := 2
	res := insertPos(arr, tar)
	if res != -1 {
		fmt.Printf("the index of the insert number is %d\n", res)
	} else {
		fmt.Println("not found")
	}
}

func insertPos(arr []int, tar int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == tar {
			return mid
		} else if arr[mid] < tar {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
