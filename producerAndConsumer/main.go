package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	resource := make(chan int, 10)
	go producer(resource)
	go consumer(resource, 1)
	// go consumer(resource, 2)
	wg.Wait()
}

func producer(resource chan<- int) {
	i := 0
	for {
		fmt.Println("produce---", i)
		resource <- i
		i++
		time.Sleep(time.Second)
	}
}

func consumer(resource <-chan int, no int) {
	for {
		select {
		case i := <-resource:
			fmt.Println("consume(", no, ")----", i)
		default:
			fmt.Println("empty---(", no, ")")
		}
		time.Sleep(time.Second)
	}
}
