package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1099511627776", 10) // 2^40
	b.SetString("2199023255552", 10) // 2^41

	result := new(big.Int).Mul(a, b)
	fmt.Printf("%s * %s = %s\n", a, b, result)

	result = new(big.Int).Div(b, a)
	fmt.Printf("%s / %s = %s\n", b, a, result)

	result = new(big.Int).Add(a, b)
	fmt.Printf("%s + %s = %s\n", a, b, result)

	result = new(big.Int).Sub(b, a)
	fmt.Printf("%s - %s = %s\n", b, a, result)
}
