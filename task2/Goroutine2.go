package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	tasks := []func(){
		func() { time.Sleep(1 * time.Second); fmt.Println("任务1完成") },
		func() { time.Sleep(2 * time.Second); fmt.Println("任务2完成") },
		func() { time.Sleep(500 * time.Millisecond); fmt.Println("任务3完成") },
	}

	for i, task := range tasks {
		wg.Add(1)
		go func(id int, f func()) {
			defer wg.Done()
			start := time.Now()
			f()
			fmt.Printf("任务%d耗时 %v\n", id+1, time.Since(start))
		}(i, task)
	}

	wg.Wait()
}
