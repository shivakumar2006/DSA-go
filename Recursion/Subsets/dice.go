package main

import "fmt"

func main() {
	dice("", 4)
}

func dice(process string, tar int) {
	if tar == 0 {
		fmt.Println(process)
		return
	}

	for i := 1; i <= 6 && i <= tar; i++ {
		dice(process+fmt.Sprint(i), tar-i)
	}
}

// func dice(process string, tar int) {
// 	if tar == 0 {
// 		fmt.Println(process)
// 		return
// 	}

// 	for i := 1; i <= 6 && i <= tar; i++ {
// 		dice(process+fmt.Sprint(i), tar-i)
// 	}
// }
