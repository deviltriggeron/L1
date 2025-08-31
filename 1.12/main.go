package main

import (
	"fmt"
	"slices"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	var str []string

	for i := range sequence {
		if !slices.Contains(str, sequence[i]) {
			str = append(str, sequence[i])
		}
	}

	fmt.Println(str)
}
