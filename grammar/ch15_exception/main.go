package main

import (
	"fmt"
)

func panic3() {
	fmt.Println("panic3 start")
	panic("panic in panic3")
	fmt.Println("panic3 end")
}

func panic2() {
	fmt.Println("panic2 start")
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	panic3()
	fmt.Println("panic2 end")
}

func panic1() {
	fmt.Println("panic1 start")
	panic2()
	fmt.Println("panic1 end")
}

func main() {

	// panic 会向上传递，直到遇到有 recover 的函数
	// recover 只能捕获以当前函数为首的调用链中的函数中的 panic
	panic1()
	/*
		panic1 start
		panic2 start
		panic3 start
		panic in panic3
		panic1 end
	*/

	// 使用 recover 捕获异常
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()
	fmt.Println("a")
	panic("exception occurs")
	fmt.Println("b") // 该代码不会被执行
}
