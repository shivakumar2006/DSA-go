// Minimum non-zero product of the array element
// https://leetcode.com/problems/minimum-non-zero-product-of-the-array-elements/description/

package main

import "fmt"

const mod int64 = 1000000007

func modPow(base, exponent, mod int64) int64 {
	base %= mod // <<< important: reduce before any squaring
	result := int64(1)
	for exponent > 0 {
		if exponent&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exponent >>= 1
	}
	return result
}

func minNonZeroProduct(p int) int {
	maxNum := (int64(1) << p) - 1
	secondMax := maxNum - 1
	exponent := (int64(1) << (p - 1)) - 1
	powPart := modPow(secondMax, exponent, mod)
	result := (powPart * maxNum) % mod
	if result < 0 {
		result += mod
	} // extra safety (not really needed)
	return int(result)
}

func main() {
	p := 3
	fmt.Println(minNonZeroProduct(p))
}
