// permutation count

package main

import "fmt"

func main() {
	count := permutationCount("", "abc")
	fmt.Println("Permutation count is : ", count)
}

func permutationCount(process, unprocess string) int {
	if len(unprocess) == 0 {
		return 1
	}

	count := 0
	ch := unprocess[0]
	rest := unprocess[1:]

	for i := 0; i <= len(process); i++ {
		newProcess := process[:i] + string(ch) + process[i:]
		count += permutationCount(newProcess, rest)
	}
	return count
}
