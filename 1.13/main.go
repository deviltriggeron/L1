package main

import "fmt"

func main() {
	n1 := 2
	n2 := 3

	fmt.Printf("Before %d %d\n", n1, n2)
	n1 ^= n2
	n2 ^= n1
	n1 ^= n2
	fmt.Printf("After %d %d", n1, n2)

}
