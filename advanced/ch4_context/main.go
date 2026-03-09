package main

import (
	"context"
	"fmt"
	"time"
)

func goroutine1(ctx context.Context, name string) {
	for {
		ch := ctx.Done()
		select {
		case <-ch:
			fmt.Printf("%s end!!\n", name)
			return
		default:
			fmt.Printf("%s is running...\n", name)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func goroutine2(ctx context.Context) {
	fmt.Printf("goroutine2: %v\n", ctx.Value("name"))
}

func main() {
	// context 用于并发控制，在不需要子 goroutine 执行的时候，可以通过 context 通知子 goroutine 优雅的关闭

	// 例 1
	// 使用 context.WithCancel 创建子context 和 cancel 函数
	// 向子 goroutine 传递子context，goroutine 通过 Done 获取仅读 channel
	// 父 goroutine 调用 cancel 函数，向 channel 发送数据
	// 子 goroutine 通过仅读 channel 接收数据后，手动返回
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	go goroutine1(ctx, "tutu1")
	go goroutine1(ctx, "tutu2")
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)

	// 例 2
	// 使用 context.WithDeadline 来创建 子context 和 cancel 函数
	// 和 context.WithCancel 不同的是，该函数可以指定 Deadline，如果到 Deadline 时，没有手动 cancel，会自动向 channel 中发送数据
	ctx2, cancel := context.WithDeadline(rootCtx, time.Now().Add(time.Second*2))
	go goroutine1(ctx2, "tutu3")
	go goroutine1(ctx2, "tutu4")
	time.Sleep(time.Second * 4)
	fmt.Printf("end without cancel(), but child goroutine is already done..\n")
	// cancel()

	// 例 3
	// 使用
	ctx3, cancel := context.WithTimeout(rootCtx, time.Second*2)
	go goroutine1(ctx3, "tutu5")
	go goroutine1(ctx3, "tutu6")
	time.Sleep(time.Second * 4)
	fmt.Printf("end without cancel(), but child goroutine is already done..\n")

	// 例 4
	// 利用 context.WithValue 向 子goroutine 传值
	ctx4 := context.WithValue(rootCtx, "name", "tutu7")
	go goroutine2(ctx4)

	time.Sleep(time.Millisecond * 500)

}
