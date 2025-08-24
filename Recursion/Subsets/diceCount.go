package main

import "fmt"

func main() {
	ans := dice("", 4)
	fmt.Println(ans)
}

func dice(process string, tar int) int {
	if tar == 0 {
		return 1
	}

	count := 0
	for i := 1; i <= 6 && i <= tar; i++ {
		count += dice(process+fmt.Sprint(i), tar-i)
	}
	return count
}
