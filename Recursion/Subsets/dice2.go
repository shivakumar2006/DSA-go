package main

import "fmt"

func main() {
	ans := dice("", 5)
	fmt.Println(ans)
}

func dice(process string, tar int) []string {
	if tar == 0 {
		return []string{process}
	}

	var result []string
	for i := 1; i <= 6 && i <= tar; i++ {
		newResult := dice(process+fmt.Sprint(i), tar-i)
		result = append(result, newResult...)
	}
	return result
}
