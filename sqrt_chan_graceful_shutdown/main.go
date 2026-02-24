package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	Input   int
	Squared int
}

func main() {
	inputChan := make(chan int)
	resultChan := make(chan Result)
	wg := &sync.WaitGroup{}
	numsWorkers := 3
	inputNumsContext, _ := context.WithTimeout(context.Background(), 3000*time.Millisecond)

	go inputNums(inputChan, inputNumsContext)

	wg.Add(numsWorkers)
	for i := 1; i <= numsWorkers; i++ {
		go sqrtChan(inputChan, resultChan, wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for val := range resultChan {
		fmt.Printf("Input: %d → Squared: %d\n", val.Input, val.Squared)
	}
}

func sqrtChan(ch chan int, outChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		outChan <- Result{
			Input:   val,
			Squared: val * val,
		}
	}
}

func inputNums(ch chan int, ctx context.Context) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Числа переданы в канал")
			return
		default:
			ch <- rand.Intn(10) + 1
		}
		time.Sleep(250 * time.Millisecond)
	}
}
