package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Start")
	Sleep(2 * time.Second)
	fmt.Println("End")
}
