package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewInt(int64(math.Pow(2, 28)))
	b := big.NewInt(int64(math.Pow(2, 25)))

	res := big.NewInt(0).Add(a, b)
	fmt.Printf("a + b = %d\n", res)

	res = big.NewInt(0).Sub(a, b)
	fmt.Printf("a - b = %d\n", res)

	res = big.NewInt(0).Mul(a, b)
	fmt.Printf("a * b = %d\n", res)

	res = big.NewInt(0).Div(a, b)
	fmt.Printf("a / b = %d\n", res)
}
