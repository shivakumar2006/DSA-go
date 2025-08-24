// letter combination of a phone number using []slice{}

package main

import "fmt"

func main() {
	ans := pad("", "12")
	fmt.Println(ans)
}

func pad(process, unprocess string) []string {
	if len(unprocess) == 0 {
		return []string{process}
	}

	digit := int(unprocess[0] - '0')
	start := (digit - 1) * 3
	end := digit * 3

	ans := []string{}
	for i := start; i < end; i++ {
		ch := string('a' + rune(i))
		newReuslt := pad(process+ch, unprocess[1:])
		ans = append(ans, newReuslt...)
	}
	return ans
}

// func pad(process, unprocess string) []string {
// 	if len(unprocess) == 0 {
// 		return []string{process}
// 	}

// 	digit := int(unprocess[0] - '0')
// 	start := (digit - 1) * 3
// 	end := digit * 3

// 	result := []string{}

// 	for i := start; i < end; i++ {
// 		ch := string('a' + rune(i))
// 		newResult := pad(process+ch, unprocess[1:])
// 		result = append(result, newResult...)
// 	}
// 	return result
// }
