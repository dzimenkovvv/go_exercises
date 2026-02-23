package main

import (
	"fmt"
	"sync"
	"time"
)

func name(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(name)
		time.Sleep(250 * time.Millisecond)
	}
}
func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go name(wg, "Bob")

	wg.Add(1)
	go name(wg, "Alice")

	wg.Add(1)
	go name(wg, "Charlie")

	wg.Wait()

	time.Sleep(3 * time.Second)
	fmt.Println("")
	fmt.Println("Main завершен!")
}
