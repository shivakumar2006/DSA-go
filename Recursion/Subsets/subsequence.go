package main

import "fmt"

func main() {
	str := "abcd"
	subseq("", str)
}

func subseq(process, unprocess string) {
	if len(unprocess) == 0 {
		fmt.Println(process)
		return
	}

	ch := string(unprocess[0])
	rest := unprocess[1:]

	// include cahracters
	subseq(process+ch, rest)

	// exclude
	subseq(process, rest)
}

// func subseq(process, unprocess string) {
// 	if len(unprocess) == 0 {
// 		fmt.Println(process)
// 		return
// 	}

// 	ch := string(unprocess[0])
// 	rest := unprocess[1:]

// 	subseq(process+ch, rest)

// 	subseq(process, rest)
// }

// func subseq(process, unprocess string) {
// 	if len(unprocess) == 0 {
// 		fmt.Println(process)
// 		return
// 	}

// 	ch := string(unprocess[0])
// 	rest := unprocess[1:]

// 	subseq(process+ch, rest)

// 	subseq(process, rest)
// }

// func subseq(process, unprocess string) {
// 	if len(unprocess) == 0 {
// 		fmt.Println(process)
// 		return
// 	}

// 	ch := string(unprocess[0])
// 	rest := unprocess[1:]

// 	subseq(process+ch, rest)

// 	subseq(process, rest)
// }

// func subseq(processed, unprocessed string) {
// 	if len(unprocessed) == 0 {
// 		fmt.Println(processed)
// 		return
// 	}

// 	ch := string(unprocessed[0]) // first char
// 	rest := unprocessed[1:]      // rest of the characters

// 	// include characters
// 	subseq(processed+ch, rest)

// 	// exclude characters
// 	subseq(processed, rest)
// }

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
