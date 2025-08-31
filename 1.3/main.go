package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, ch chan int) {
	for i := range ch {
		fmt.Printf("Worker %d: got job %d\n", id, i)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	ch := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, ch)
	}

	jobID := 1
	for {
		ch <- jobID
		jobID++
		time.Sleep(500 * time.Millisecond)
	}
}
