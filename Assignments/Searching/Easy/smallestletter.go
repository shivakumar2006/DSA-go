// Find smallest letter greater than target
// https://leetcode.com/problems/find-smallest-letter-greater-than-target/description/

package main

import "fmt"

func main() {
	input := []byte{'c', 'j', 'w'}
	var target byte = 'a'
	fmt.Println("the smalles character greater than a is : ", string(nextGreatestLetter(input, target)))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	start := 0
	end := len(letters) - 1
	for start <= end {
		mid := start + (end-start)/2
		if letters[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return letters[start%len(letters)]
}

// func nextGreatestLetter(letters []byte, target byte) byte {
// 	start := 0
// 	end := len(letters) - 1

// 	for start <= end {
// 		mid := start + (end-start)/2
// 		if letters[mid] <= target {
// 			start = mid + 1
// 		} else {
// 			end = mid - 1
// 		}
// 	}
// 	return letters[start%len(letters)]
// }
