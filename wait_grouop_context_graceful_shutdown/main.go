package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func numsInChannel(ctx context.Context, jobs chan int, wg *sync.WaitGroup, n int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", n, "end!")
			return
		case jobs <- rand.Intn(10):
			time.Sleep(250 * time.Millisecond)
		}
	}
}
func main() {
	jobs := make(chan int, 10)
	numWorkers := 3
	wg := &sync.WaitGroup{}
	parentContext, _ := context.WithTimeout(context.Background(), 2*time.Second)

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go numsInChannel(parentContext, jobs, wg, i)
	}

	go func() {
		wg.Wait()
		close(jobs)
	}()

	for val := range jobs {
		fmt.Println(val)
	}
}
