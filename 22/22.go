package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2 << 22)
	b := big.NewInt(35 << 22)
	result := new(big.Int)

	result.Add(a, b)
	fmt.Println(result)

	result.Div(b, a)
	fmt.Println(result)

	result.Sub(b, a)
	fmt.Println(result)

	result.Mul(b, a)
	fmt.Println(result)
}
