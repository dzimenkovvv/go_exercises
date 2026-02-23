package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func numsInChannel(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	var sum int
	numWorkers := 3

	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go numsInChannel(ch, wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Print(val, " ")
		sum += val
	}
	fmt.Println("\nTotal sum:", sum)
}
