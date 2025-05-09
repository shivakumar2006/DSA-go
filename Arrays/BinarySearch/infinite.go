// find the target element in a infinite sorted array...
// this is the sorted array so, obviously we use binary search...
// this is amazon interview question...

package main

import (
	"fmt"
)

var arr = []int{2, 4, 5, 7, 9, 11, 13, 16, 20, 25, 30, 35, 40}

func main() {
	tar := 25
	res := search(tar)
	if res != -1 {
		fmt.Printf("Target %d found at the index : %d\n", tar, res)
	} else {
		fmt.Println("Target not found...")
	}
}

func search(tar int) int {
	l := 0
	r := 1

	for get(r) < tar {
		l = r
		r *= 2
	}

	for l <= r {
		mid := l + (r-l)/2
		val := get(mid)
		if val == tar {
			return mid
		} else if val < tar {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func get(res int) int {
	if res >= len(arr) {
		return int(^uint(0) >> 1)
	}
	return arr[res]
}
