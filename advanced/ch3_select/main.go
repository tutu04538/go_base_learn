package main

import (
	"fmt"
	"time"
)

func main() {
	// select 是 Go 语言层面提供的一种 IO 多路复用机制，用于检测当前 goroutine 连接的多个 channel 是否有数据准备完毕，可以用于读和写
	// select 的 case 必须是 channel 的读写操作
	// 当有多个 case 同时满足时，会随机选择一个
	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	ch1 <- 1
	ch2 <- 2
	select {
	case v := <-ch1:
		fmt.Printf("ch1: %d\n", v)
	case v := <-ch2:
		fmt.Printf("ch2: %d\n", v)
	default:
		fmt.Printf("no ch\n")
	}

	ch3, ch4 := make(chan int, 1), make(chan int, 1)
	// 当注释掉下面的代码时，程序会 panic，因为会陷入永久阻塞
	// 下面的代码正常运行时，select 会阻塞直到 goroutine 执行完毕
	go func() {
		time.Sleep(1000 * time.Millisecond)
		ch3 <- 1
	}()

	select {
	case v := <-ch3:
		fmt.Printf("ch3: %d\n", v)
	case v := <-ch4:
		fmt.Printf("ch4: %d\n", v)
	}

	time.Sleep(1000 * time.Millisecond)
}
