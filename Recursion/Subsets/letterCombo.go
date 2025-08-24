// letter combination of a phone number...

package main

import "fmt"

func main() {
	pad("", "12")
}

func pad(process, unprocess string) {
	if len(unprocess) == 0 {
		fmt.Println(process)
		return
	}

	digit := int(unprocess[0] - '0')
	start := (digit - 1) * 3
	end := digit * 3
	for i := start; i < end; i++ {
		ch := string('a' + rune(i))
		pad(process+ch, unprocess[1:])
	}
}

// solve same question using slice
// func pad(process, unprocess string) []string {
// 	if len(unprocess) == 0 {
// 		return []string{process}
// 	}

// 	digit := int(unprocess[0] - '0')
// 	start := (digit - 1) * 3
// 	end := digit * 3

// 	ans := []string{}

// 	for i := start; i < end; i++ {
// 		ch := string('a' + rune(i))
// 		subResult := pad(process+ch, unprocess[1:])
// 		ans = append(ans, subResult...)
// 	}
// 	return ans
// }

// func pad(process, unprocess string) {
// 	if len(unprocess) == 0 {
// 		fmt.Println(process)
// 		return
// 	}

// 	digit := int(unprocess[0] - '0')
// 	start := (digit - 1) * 3
// 	end := digit * 3
// 	for i := start; i < end; i++ {
// 		ch := string('a' + rune(i))
// 		pad(process+ch, unprocess[1:])
// 	}
// }

// func pad(process, unprocess string) {
// 	if len(unprocess) == 0 {
// 		fmt.Println(process)
// 		return
// 	}

// 	digit := int(unprocess[0] - '0') // convert char to int
// 	start := (digit - 1) * 3         // start index
// 	end := digit * 3                 // end index

// 	for i := start; i < end; i++ {
// 		ch := string('a' + rune(i)) // convert int to rune then to string
// 		pad(process+ch, unprocess[1:])
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	combo("", "23")
// }

// func combo(process, unprocess string) {
// 	if len(unprocess) == 0 {
// 		fmt.Println(process)
// 		return
// 	}

// 	digit := unprocess[0] - '0'

// 	// Skip digits like 0 or 1 since they have no letters
// 	if digit < 2 || digit > 9 {
// 		combo(process, unprocess[1:])
// 		return
// 	}

// 	start := (digit - 2) * 3
// 	end := start + 3

// 	// handle 7, 8, 9extra letters
// 	if digit == 7 {
// 		end++
// 	} else if digit == 8 {
// 		start += 1
// 		end += 1
// 	} else if digit == 9 {
// 		start += 1
// 		end += 2
// 	}

// 	for i := start; i < end; i++ {
// 		ch := string('a' + i)
// 		combo(process+ch, unprocess[1:])
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	pad("", "12")
// }

// func pad(process, unprocess string) {
// 	if len(unprocess) == 0 {
// 		fmt.Println(process)
// 		return
// 	}

// 	digit := unprocess[0] - '0'
// 	rest := unprocess[1:]
// 	for i := (digit - 1) * 3; i < digit * 3; i++ {
// 		ch := ('a' + 1)
// 		pad(process+ch, rest)
// 	}
// }
