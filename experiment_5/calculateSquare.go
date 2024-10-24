package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 500; x++ {
		out <- x
	}
	close(out) //关闭通道
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	i := 0
	for v := range in {
		fmt.Println(i, v)
		i++
	}
}
func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
