package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go produce1(ch1)
	go produce2(ch2)

	go consume(ch1, ch2)

	time.Sleep(1e9)
}
func produce1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}
func produce2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func consume(ch1, ch2 chan int) {
	// 死循环
	for {
		select {
		case v := <-ch1:
			fmt.Printf("channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("channel 2: %d\n", v)
		}
	}
}
