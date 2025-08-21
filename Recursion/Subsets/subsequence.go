package main

import "fmt"

func main() {
	str := "abc"
	subseq("", str)
}

func subseq(processed, unprocessed string) {
	if len(unprocessed) == 0 {
		fmt.Println(processed)
		return
	}

	ch := string(unprocessed[0]) // first char
	rest := unprocessed[1:]      // rest of the characters

	// include characters
	subseq(processed+ch, rest)

	// exclude characters
	subseq(processed, rest)
}

// func subsequences(processed, unprocessed string) {
// 	if len(unprocessed) == 0 {
// 		fmt.Println(processed)
// 		return
// 	}

// 	ch := string(unprocessed[0])

// 	// include characters
// 	subsequences(processed+ch, unprocessed[1:])

// 	// exclude characters
// 	subsequences(processed, unprocessed[1:])
// }
