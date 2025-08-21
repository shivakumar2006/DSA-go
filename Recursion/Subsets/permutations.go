package main

import "fmt"

func main() {
	permutations("", "abc")
}

func permutations(process string, unprocess string) {
	if len(unprocess) == 0 {
		fmt.Println(process)
		return
	}

	ch := unprocess[0]
	rest := unprocess[1:]

	// insert character in every possible position of the process
	for i := 0; i <= len(process); i++ {
		newProcess := process[:i] + string(ch) + process[i:]
		permutations(newProcess, rest)
	}
}
