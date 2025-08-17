package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 6, 69, 99}
	fmt.Println("Is the target is find : ", find(arr, 69, 0))
	fmt.Println("the index of the target element is : ", findIndex(arr, 69, 0))
	fmt.Println("the target index is from the last index is : ", findLastIndex(arr, 69, len(arr)-1))
	fmt.Println(findAllIndex(arr, 69, 0))
}

func find(arr []int, tar int, index int) bool {
	if index == len(arr)-1 {
		return false
	}

	return arr[index] == tar || find(arr, tar, index+1)
}

func findIndex(arr []int, tar int, index int) int {
	if index == len(arr) {
		return -1
	}

	if arr[index] == tar {
		return index
	} else {
		return findIndex(arr, tar, index+1)
	}
}

func findLastIndex(arr []int, tar int, index int) int {
	if index == -1 {
		return -1
	}

	if arr[index] == tar {
		return index
	} else {
		return findLastIndex(arr, tar, index-1)
	}
}

func findAllIndex(arr []int, tar int, index int) []int {
	if index == len(arr)-1 {
		return []int{}
	}

	result := findAllIndex(arr, tar, index+1)

	if arr[index] == tar {
		result = append([]int{index}, result...) // prepend current index
	}
	return result
}
