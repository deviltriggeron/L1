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
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	jobID := 1
loop:
	for {
		select {
		case <-signalChan:
			fmt.Println("\nReceived interrupt signal, shutting down...")
			signal.Stop(signalChan)
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
