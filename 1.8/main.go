package main

import "fmt"

func setBit(n int64, i int, b int) int64 {
	if b == 1 {
		n |= 1 << i
	} else if b == 0 {
		n &^= 1 << i
	}
	return n
}

func main() {
	var n int64 = 5

	res := setBit(n, 0, 0)

	fmt.Printf("(%04b)", n)
	fmt.Printf("(%04b)", res)

	fmt.Printf("\n %d %d", n, res)
}
