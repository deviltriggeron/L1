package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range ch {
		fmt.Printf("Worker %d: got job %d\n", id, job)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <num_workers> <duration seconds>")
		return
	}

	parseArg := func(arg string) int {
		val, err := strconv.Atoi(arg)
		if err != nil || val <= 0 {
			fmt.Printf("Not number\n")
			os.Exit(1)
		}
		return val
	}

	numWorkers := parseArg(os.Args[1])
	durationSeconds := parseArg(os.Args[2])

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	jobID := 1
	timeout := time.After(time.Duration(durationSeconds) * time.Second)

loop:
	for {
		select {
		case <-signalChan:
			fmt.Println("\nReceived interrupt signal, shutting down...")
			signal.Stop(signalChan)
			close(ch)
			break loop
		case <-timeout:
			fmt.Println("\nTime is up, shutting down...")
			close(ch)
			break loop
		case ch <- jobID:
			jobID++
			time.Sleep(500 * time.Millisecond)
		}

	}

	wg.Wait()
	fmt.Println("All workers stopped. Exiting.")
}
