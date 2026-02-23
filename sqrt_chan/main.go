package main

import (
	"fmt"
	"math/rand"
	"sync"
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

	go func() {
		defer close(inputChan)
		for i := 1; i <= 10; i++ {
			inputChan <- rand.Intn(10) + 1
		}
	}()

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
