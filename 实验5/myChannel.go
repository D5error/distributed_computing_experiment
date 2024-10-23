package main

import (
	"fmt"
	"time"
)

func main() {
	// chan TYPE: 一种内部为TYPE类型的channel类型
	// chan <- TYPE:只能发送TYPE类型
	// <- chan TYPE: 只能接收TYPE类型
	nums := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go sum(nums[:len(nums)/2], c)
	go sum(nums[len(nums)/2:], c)
	x, y := <-c, <-c
	fmt.Println("x", x, "y", y, "x+y", x+y)

	go test(1, 10, c)
	go test(1, 300, c)
	t1, t2 := <-c, <-c
	fmt.Println("1 + .... =", t1, "and ", t2)
}

func sum(nums []int, c chan int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	c <- sum
}

func test(a int, b int, c chan int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
		time.Sleep(10500 * time.Microsecond)
		fmt.Println("a:", a, "b:", b, "i:", i)
	}
	fmt.Println("c:", c)
	c <- sum
	// close(c)
}
