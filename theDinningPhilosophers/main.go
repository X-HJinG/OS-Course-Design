package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	N int = 5 //哲学家人数与筷子数量
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	var chopsticks [N]chan int
	status := make(chan int, N-1)
	//初始化筷子数量
	for i := 0; i < N; i++ {
		chopsticks[i] = make(chan int, 1)
		chopsticks[i] <- 1
	}
	//初始化限制进餐人数N-1
	for i := 0; i < N-1; i++ {
		status <- 1
	}
	for i := 0; i < N; i++ {
		go philosopher(i, chopsticks, status)
	}
	wg.Wait()
}

func philosopher(no int, chopsticks [N]chan int, status chan int) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Int()%(50+1)))
		//进入等待
		<-status
		//拿起右边筷子
		<-chopsticks[no]
		//拿起左边筷子
		<-chopsticks[(no+1)%N]
		fmt.Println("Philosopher:", no, " are eating.")
		//放下左边筷子
		chopsticks[(no+1)%N] <- 1
		//放下右边筷子
		chopsticks[no] <- 1
		//退出等待
		status <- 1
	}
}
