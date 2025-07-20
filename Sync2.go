package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1)

	// 启动110个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程递增计数器1000次
			for j := 0; j < 1000; j++ {
				ch <- struct{}{}
				counter++
				<-ch
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter)
}
