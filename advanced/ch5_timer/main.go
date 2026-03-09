package main

import (
	"fmt"
	"time"
)

// time.Newtimer 创建一个 timer
// timer.C 是一个仅读 channel
// 到达时间时 timer 会向其中写数据
func startASimpleTimer() {
	fmt.Println("startASimpleTimer start")
	timer := time.NewTimer(time.Second)
	<-timer.C
	fmt.Println("startASimpleTimer end after one second")
}

// 使用 timer.Stop() 结束一个 timer
// 使用 timer.Reset() 重置一个 timer
func endASimpleTimer() {
	fmt.Println("endASimpleTimer start")
	timer := time.NewTimer(time.Second * 5)
	ok := timer.Stop()
	if ok {
		fmt.Println("timer end before time exceeds")
	}

	timer.Reset(time.Millisecond * 500)
	<-timer.C
	fmt.Println("timer Reset")
}

func AfterFuncTimer() {
	simpleFunc := func() {
		fmt.Println("AfterFunc Timer end")
	}
	timer := time.AfterFunc(500*time.Millisecond, simpleFunc)
	defer timer.Stop()
	time.Sleep(time.Second)
}

func AfterTimer() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("goroutine end")
	case <-time.After(time.Second):
		fmt.Println("goroutine doesn't end, but time already exceeds!!")
	}

}

func simpleTicker() chan struct{} {
	ticker := time.NewTicker(time.Millisecond * 500)
	ch := make(chan struct{})

	go func(ticker *time.Ticker) {
		defer ticker.Stop() // 不会关闭 ticker.C 这个管道
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker acts!")
			case <-ch:
				fmt.Println("ticker end from parent goroutine...")
				return
			}
		}
	}(ticker)

	return ch
}

func tickerTimer() {

	ch := simpleTicker()

	time.Sleep(time.Second * 2)
	ch <- struct{}{}
	close(ch)

	time.Sleep(time.Millisecond * 500)
}

func main() {
	// startASimpleTimer()
	// endASimpleTimer()
	// AfterFuncTimer()
	// AfterTimer()
	tickerTimer()
}
