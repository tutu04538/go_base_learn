package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func add(ch chan bool, num *int) {
	ch <- true
	*num = *num + 1
	<-ch
}

func main() {
	// channel 是一个可以收发数据的管道
	// var channel_name chan channel_type
	// var channel_name [size]chan channel_type 带缓存
	// channel_name := make(chan channel_type)
	// channel_name := make(chan channel_type, size)

	// ch := make(chan int)
	// ch <- v 		向管道 ch 中发送数据 v
	// v := <-ch	从管道 ch 中读取数据
	// close(ch)	关闭管道

	// 关闭管道后仍可以从管道中接收数据
	ch := make(chan int, 6)
	ch <- 1
	close(ch)
	go func() {
		for range 6 {
			v := <-ch
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second)

	// 判定句式读取
	ch2 := make(chan int, 6)
	ch2 <- 2
	close(ch2)
	// 如果不加 close(ch2)，下面的 v, ok := <-ch2 会阻塞，直到主进程结束
	go func() {
		for range 6 {
			if v, ok := <-ch2; ok {
				fmt.Println(v)
			} else {
				fmt.Printf("ch2 is empty, v=%d\n", v)
			}
		}
	}()

	time.Sleep(time.Second)
	/*
		2
		ch2 is empty, v=0
		ch2 is empty, v=0
		ch2 is empty, v=0
		ch2 is empty, v=0
		ch2 is empty, v=0
	*/

	// for range 读取，可以一直读，直到另一段关闭这个 channel
	ch3 := make(chan int, 6)
	ch3 <- 1
	ch3 <- 2
	close(ch3)
	go func() {
		for v := range ch3 {
			fmt.Printf("read from ch3: %d\n", v)
		}
	}()

	time.Sleep(time.Second)

	// 单向 channel
	type RChannel <-chan int
	type SChannel chan<- int

	ch4 := make(chan int)

	go func() {
		var sch SChannel = ch4
		sch <- 114514
	}()

	go func() {
		var rch RChannel = ch4
		v := <-rch
		fmt.Printf("read from rch: %d\n", v)
	}()

	time.Sleep(time.Second)

	// 协程之间通过 channel 传递数据
	a := []int{1, 2, 3, 4, 5, 6}
	ch5 := make(chan int) // channel 在使用之前必须 make，否则发送和接收数据都会阻塞
	go func() {
		sum(a[:len(a)/2], ch5)
	}()

	go sum(a[len(a)/2:], ch5)

	b, c := <-ch5, <-ch5
	fmt.Printf("b=%d, c=%d\n", b, c)

	time.Sleep(time.Second)

	// 有缓冲 channel 和无缓冲 channel 内部都会有 lock 来控制并发访问
	// 无缓冲 channel 在接收到数据后，如果没有消费者，再次写入时就会阻塞
	// 有缓冲 channel 在队列满时，如果没有消费者，再次写入也会阻塞

	// channel 是并发安全的，多个协程同时读取 channel 中的数据，不会产生并发安全问题
	// 利用 channel 实现锁操作
	ch6 := make(chan bool, 1)

	var num int
	for range 100 {
		go add(ch6, &num)
	}
	fmt.Println(num)
}
