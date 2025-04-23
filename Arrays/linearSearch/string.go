package main

func search(str string, target string) int {
	for i := 0; i <= len(str); i++ {
		if string(str[i]) == target {
			return i
		}
	}
	return -1
}
