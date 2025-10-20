// K-th Symbol in Grammar
// http://leetcode.com/problems/k-th-symbol-in-grammar/description/

package main

import "fmt"

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	// 2^n-2
	mid := 1 << (n - 2) // middle position of current row
	if k <= mid {
		return kthGrammar(n-1, k) // left half: same as previous row
	} else {
		return 1 - kthGrammar(n-1, k-mid) // right half: flip previous row value
	}
}

func main() {
	n := 2
	k := 1
	fmt.Println(kthGrammar(n, k))
}
