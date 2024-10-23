package main

import (
	"fmt"
	"sync"
)

func processTask(task int, i int) {
	fmt.Println("Processing task", task, "channel", i)
}

func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// 定义并发数为3的批量处理函数
	batchSize := 3
	var wg sync.WaitGroup
	taskChan := make(chan int)
	for i := 0; i < batchSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for task := range taskChan {
				processTask(task, i)
			}
		}(i)
	}
	// 创建了3个ProcessTask子线程，它们都使用了同一个channel

	// 将任务分发到taskChan通道中
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)
	wg.Wait()
}
