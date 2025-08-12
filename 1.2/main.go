package main

import (
	"fmt"
	"time"
)

func pow(a int) {
	fmt.Println(a * a)
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	for i := range arr {
		go pow(i)
	}
	time.Sleep(time.Second)
}
