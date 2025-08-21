// permutation using slice

package main

import "fmt"

func main() {
	result := permutations([]rune{}, []rune("abc"))
	for _, permu := range result {
		fmt.Println(string(permu))
	}
}

func permutations(process, unprocess []rune) [][]rune {
	if len(unprocess) == 0 {
		return [][]rune{append([]rune{}, process...)}
	}

	var result [][]rune
	ch := unprocess[0]
	rest := unprocess[1:]

	for i := 0; i <= len(process); i++ {
		newProcess := make([]rune, len(process)+1)
		copy(newProcess[:i], process[:i])
		newProcess[i] = ch
		copy(newProcess[i+1:], process[i:])
		result = append(result, permutations(newProcess, rest)...)
	}
	return result
}

// func permutations(process, unprocess []rune) [][]rune {
// 	if len(unprocess) == 0 {
// 		return [][]rune{append([]rune{}, process...)}
// 	}

// 	var result [][]rune
// 	ch := unprocess[0]
// 	rest := unprocess[1:]

// 	for i := 0; i <= len(process); i++ {
// 		// create new slice with character inserted at index i
// 		newProcess := make([]rune, len(process)+1)
// 		copy(newProcess[:i], process[:i])
// 		newProcess[i] = ch
// 		copy(newProcess[i+1:], process[i:])
// 		// collects all results recursively
// 		result = append(result, permutations(newProcess, rest)...)
// 	}
// 	return result
// }
