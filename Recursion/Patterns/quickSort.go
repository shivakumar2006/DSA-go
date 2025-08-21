package main

import "fmt"

func main() {
	arr := []int{5, 4, 3, 2, 1}
	sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func sort(arr []int, low, high int) {
	if low >= high {
		return
	}

	start := low
	end := high
	mid := start + (end-start)/2
	pivot := arr[mid]

	for start <= end {
		for arr[start] < pivot {
			start++
		}
		for arr[end] > pivot {
			end--
		}

		if start <= end {
			temp := arr[start]
			arr[start] = arr[end]
			arr[end] = temp
			start++
			end--
		}
	}

	sort(arr, low, end)
	sort(arr, start, high)
}

// func sort(arr []int, low, high int) {
// 	if low >= high {
// 		return
// 	}

// 	start := low
// 	end := high
// 	mid := start + (end-start)/2
// 	pivot := arr[mid]

// 	for start <= end {
// 		for arr[start] < pivot {
// 			start++
// 		}
// 		for arr[end] > pivot {
// 			end--
// 		}

// 		if start <= end {
// 			temp := arr[start]
// 			arr[start] = arr[end]
// 			arr[end] = temp
// 			start++
// 			end--
// 		}
// 	}

// 	sort(arr, low, end)
// 	sort(arr, start, high)
// }

// func sort(arr []int, low, high int) {
// 	if low >= high {
// 		return
// 	}

// 	start := low
// 	end := high
// 	mid := start + (end-start)/10
// 	pivot := arr[mid]

// 	for start <= end {
// 		for arr[start] < pivot {
// 			start++
// 		}
// 		for arr[end] > pivot {
// 			end--
// 		}

// 		if start <= end {
// 			temp := arr[start]
// 			arr[start] = arr[end]
// 			arr[end] = temp
// 			start++
// 			end--
// 		}
// 	}
// 	sort(arr, low, end)
// 	sort(arr, start, high)
// }

// func sort(arr []int, low int, high int) {
// 	if low >= high {
// 		return
// 	}

// 	start := low
// 	end := high
// 	mid := start + (end-start)/2
// 	pivot := arr[mid]

// 	for start <= end {
// 		for arr[start] < pivot {
// 			start++
// 		}
// 		for arr[end] > pivot {
// 			end--
// 		}

// 		if start <= end {
// 			temp := arr[start]
// 			arr[start] = arr[end]
// 			arr[end] = temp
// 			start++
// 			end--
// 		}
// 	}
// 	sort(arr, low, end)
// 	sort(arr, start, high)
// }
