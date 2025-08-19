package main

import "fmt"

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	i := start
	j := mid + 1

	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			i++
		} else {
			temp := arr[j]

			for k := j; k > i; k-- {
				arr[k] = arr[k-1]
			}

			arr[i] = temp

			i++
			mid++
			j++
		}
	}
}
