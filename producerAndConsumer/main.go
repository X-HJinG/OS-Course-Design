package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	resource := make(chan int, 5)
	NewProducerPool(3, resource) //开启n个生产者
	NewConsumerPool(2, resource) //开启n个消费者
	wg.Wait()
}

func NewProducerPool(num int, resource chan<- int) {
	for i := 0; i < num; i++ {
		go producer(i, resource)
	}
}

func NewConsumerPool(num int, resource <-chan int) {
	for i := 0; i < num; i++ {
		go consumer(i, resource)
	}
}

func producer(no int, resource chan<- int) {
	for {
		i := rand.Intn(100)
		resource <- i
		fmt.Println("Producer:", no, "\tproduce:", i, "\tlen:", len(resource))
		time.Sleep(time.Second)
	}
}

func consumer(no int, resource <-chan int) {
	for v := range resource {
		time.Sleep(time.Second)
		fmt.Println("Consumer:", no, "\tconsumer:", v, "\tlen:", len(resource))
	}
}
