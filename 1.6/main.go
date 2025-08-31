package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func workerStopEndWorkers(workersCount int) {
	for i := range workersCount {
		fmt.Println("Work", i+1)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("End workers")
}

func workerStopAfterTime(duration int) {
	jobID := 1
	timeout := time.After(time.Duration(duration) * time.Second)

loop:
	for {
		select {
		default:
			fmt.Println("Work", jobID)
			jobID++
			time.Sleep(500 * time.Millisecond)
		case <-timeout:
			fmt.Println("\nTime is up, shutting down...")
			break loop
		}
	}
}

func workerStopAfterSignal(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range ch {
		fmt.Printf("Worker %d: got job %d\n", id, job)
	}
}

func workerStopContext(ctx context.Context) {
	jobID := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped by context")
			return
		default:
			fmt.Println("Work", jobID)
			time.Sleep(500 * time.Millisecond)
		}
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

	go workerStopEndWorkers(numWorkers)
	go workerStopAfterTime(durationSeconds)

	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go workerStopAfterSignal(i, ch, &wg)
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	go workerStopContext(ctx)

	<-ctx.Done()

	wg.Wait()
	time.Sleep(1 * time.Second)
}
