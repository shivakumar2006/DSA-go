package main

import "fmt"

func main() {
	arr := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(speed(arr, h))
}

func speed(arr []int, h int) int {
	start := 1
	end := findMax(arr)

	for start < end {
		mid := start + (end-start)/2

		if finish(arr, mid, h) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}

func finish(arr []int, speed, h int) bool {
	totalHours := 0
	for _, i := range arr {
		totalHours += (i + speed - 1) / speed
	}
	return totalHours <= h
}

func findMax(arr []int) int {
	max := arr[0]

	for _, i := range arr {
		if i > max {
			max = i
		}
	}
	return max
}
