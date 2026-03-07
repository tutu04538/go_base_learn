package main

import (
	"fmt"
	"sync"
	"time"
)

func myGoroutine() {
	fmt.Println("myGoroutine")
}

func myGoroutine2(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Printf("myGoroutine2: %s\n", name)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// main 是主协程
	// 通过 go 关键字开启多个协程
	go myGoroutine()
	fmt.Println("main end")
	time.Sleep(time.Second)

	// 使用 sync.WaitGroup 来等待所有协程退出
	var wg sync.WaitGroup
	wg.Add(2)

	go myGoroutine2("tutu1", &wg)
	go myGoroutine2("tutu2", &wg)

	wg.Wait()
}
