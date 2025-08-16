package main

import "fmt"

func main() {
	fun(5)
}

func fun(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	n--
	fun(n)
}
