package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex) // pointer need to new
	wg = new(sync.WaitGroup)

	wg.Add(7)
	go writeData(1)
	go readData(2)
	go readData(4)
	go readData(5)
	go readData(6)
	go readData(7)
	go writeData(3)

	wg.Wait()
	fmt.Println("main thread is over")
}

func writeData(i int) {
	defer wg.Done()
	fmt.Println("try to write:", i)
	rwMutex.Lock()
	fmt.Println("writing:", i)
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	fmt.Println("written:", i)
}

func readData(i int) {
	defer wg.Done()
	fmt.Println("try to read:", i)
	rwMutex.RLock()
	fmt.Println("reading:", i)
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock()
	fmt.Println("have read:", i)
}
