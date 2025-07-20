package main

import (
	"fmt"
	"time"
)

func receive(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

func send(ch chan<- int) {
	for i := 0; i <= 100; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

func main() {
	// 创建一个带缓冲的channel
	ch := make(chan int, 10)

	go sendOnly(ch)

	//
	go receiveOnly(ch)

	for {
		timeOut := time.After(1 * time.Second)
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeOut:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
