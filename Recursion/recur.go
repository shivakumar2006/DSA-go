package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 69, 88}
	tar := 69
	fmt.Println(search(arr, tar, 0, len(arr)-1))
}

func search(arr []int, tar int, start, end int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2
	if tar == arr[mid] {
		return mid
	} else if tar < arr[mid] {
		return search(arr, tar, start, mid-1)
	}
	return search(arr, tar, mid+1, end)
}

// package main

// import "fmt"

// func main() {
// 	ans := fibo(7)
// 	fmt.Println(ans)
// }

// func fibo(n int) int {
// 	// base condition
// 	if n < 2 {
// 		return n
// 	}

// 	return fibo(n-1) + fibo(n-2)
// }

// func main() {
// 	print(1)
// }

// func print(n int) {
// 	// base
// 	if n == 5 {
// 		fmt.Println(5)
// 		return
// 	}

// 	// body
// 	fmt.Println(n)

// 	// recursive call
// 	print(n + 1)
// }
