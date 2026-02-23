package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	intChan := make(chan int)
	strChan := make(chan string)

	go msgInt(intChan)
	go msgStr(strChan)

	for {
		if intChan == nil && strChan == nil {
			break
		}
		select {
		case val, ok := <-intChan:
			if !ok {
				intChan = nil
				continue
			}
			fmt.Println(val)
		case str, ok := <-strChan:
			if !ok {
				strChan = nil
				continue
			}
			fmt.Println(str)
		}
	}
}

func msgInt(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- rand.Intn(10)
		time.Sleep(250 * time.Millisecond)
	}
	close(ch)
}

func msgStr(ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- "hello"
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
}
