package main

import "fmt"

func main() {
	str := "abc"
	subsequences("", str)
}

func subsequences(processed, unprocessed string) {
	if len(unprocessed) == 0 {
		fmt.Println(processed)
		return
	}

	ch := string(unprocessed[0])

	// include characters
	subsequences(processed+ch, unprocessed[1:])

	// exclude characters
	subsequences(processed, unprocessed[1:])
}
