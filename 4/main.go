package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, c <-chan int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v := <-c:
			fmt.Printf("Worker %v get: %v\n", id, v)
		case <-ctx.Done():
			fmt.Printf("Worker %v done\n", id)
			return
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Enter the number of workers: ")
	var n int
	fmt.Scan(&n)

	channel := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	os_channel := make(chan os.Signal, 1)
	signal.Notify(os_channel, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i+1, channel, ctx, &wg)
	}

	for {
		select {
		case <-os_channel:
			fmt.Printf(" Shutting down...\n")
			cancel()
			wg.Wait()
			return
		case <-time.After(time.Second):
			channel <- rand.Int()
		}

	}

}
