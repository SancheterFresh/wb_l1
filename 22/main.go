package main

import (
	"fmt"
	"math/big"
)

func main() {

	var n string

	a := new(big.Float)
	fmt.Printf("Enter first value: ")
	fmt.Scan(&n)
	a.SetString(n)

	b := new(big.Float)
	fmt.Printf("Enter second value: ")
	fmt.Scan(&n)
	b.SetString(n)

	addition := new(big.Float).Add(a, b)
	multiplication := new(big.Float).Mul(a, b)
	subtraction := new(big.Float).Sub(a, b)
	division := new(big.Float).Quo(a, b)

	fmt.Printf("%1.f + %1.f = %1.f\n", a, b, addition)
	fmt.Printf("%1.f * %1.f = %1.f\n", a, b, multiplication)
	fmt.Printf("%1.f - %1.f = %1.f\n", a, b, subtraction)
	fmt.Printf("%1.f / %1.f = %1.f\n", a, b, division)
}
