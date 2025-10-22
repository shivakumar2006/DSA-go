// Minimum non-zero product of the array element
// https://leetcode.com/problems/minimum-non-zero-product-of-the-array-elements/description/

package main

import "fmt"

const mod int64 = 1e9 + 7

func modPow(secondMax, exp, mod int64) int64 {
	result := int64(1)
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * secondMax) % mod
		}
		secondMax = (secondMax * secondMax) % mod
		exp /= 2
	}
	return result
}

func minNonZeroProduct(p int) int {
	maxNum := int64((1 << p) - 1)
	secondMax := maxNum - 1
	exponent := int64((1 << (p - 1)) - 1)
	powPart := modPow(secondMax, exponent, mod)
	return int((powPart * maxNum) % mod)
}

func main() {
	p := 3
	fmt.Println(minNonZeroProduct(p))
}
