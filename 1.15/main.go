package main

import "fmt"

var justString string

func createHugeString(size int) string {
	b := make([]byte, size)
	for i := 0; i < size; i++ {
		b[i] = 'a'
	}
	return string(b)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println(len(justString))
}

func main() {
	someFunc()
}
