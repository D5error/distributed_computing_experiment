package main

import (
	"fmt"
	"time"
)

func main() {
	go say("a") // 变为异步进行
	say("b")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2000 * time.Microsecond)
		fmt.Println(s, time.Now().Format("15:04:05.000000"))
	}
}
