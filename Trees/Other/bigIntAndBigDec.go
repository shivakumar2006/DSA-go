// BigInteger and BigDecimal

// BigInteger in go *big.Int

package main

import (
	"fmt"
	"math/big"
)

func main() {

	//using big.Float
	a := big.NewFloat(0.1)
	b := big.NewFloat(0.2)

	sum := new(big.Float).Add(a, b)
	sum.SetPrec(200) // optional: 200 bit of precision
	fmt.Println(sum)

	// // using big.Rat
	// x := big.NewRat(1, 3) // 1/3
	// y := big.NewRat(2, 5) // 2/5

	// sum := new(big.Rat).Add(x, y)
	// fmt.Println("1/3 + 2/5 : ", sum) //prints fraction form: 11/15

	// // create big integer
	// a := big.NewInt(1_000_000_000_000_000_000)
	// b := big.NewInt(999_999_999_999_999_999)

	// // Addition
	// sum := new(big.Int).Add(a, b)
	// fmt.Println("Sum: ", sum)

	// // Multiply
	// product := new(big.Int).Mul(a, b)
	// fmt.Println("Product: ", product)

	// // Power
	// pow := new(big.Int).Exp(big.NewInt(2), big.NewInt(100), nil)
	// fmt.Println("2^100 = ", pow)
}
