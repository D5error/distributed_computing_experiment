package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("===run without mutex===")
	runWithoutMutex()
	fmt.Println("=======================\n")

	fmt.Println("===run with mutex===")
	runWithMutex()
	fmt.Println("====================")

}

func runWithoutMutex() {
	var count = 0
	var wg sync.WaitGroup
	var n = 10
	
	wg.Add(n) // wait group amount is n
	for i:= 0; i < n; i++ { // create 10 goroutine
		go func() { // anonymous funciton
			defer wg.Done() // delay to run this statement
			for j:= 0; j < 10000; j++ {
				count++;
			}
		}()
	}
	wg.Wait() // wait for wg's number is 0
	fmt.Println("count:", count);
}

func runWithMutex() {
	var count = 0
	var wg sync.WaitGroup
	var n = 10
	var mu sync.Mutex

	wg.Add(n) // wait group amount is n
	for i:= 0; i < n; i++ { // create 10 goroutine
		go func() { // anonymous funciton
			defer wg.Done() // delay to run this statement
			for j:= 0; j < 10000; j++ {
				mu.Lock() // try to protect the following statement
				count++
				mu.Unlock() // try to protect the above statement
			}
		}()
	}
	wg.Wait() // wait for wg's number is 0
	fmt.Println("count:", count);
}