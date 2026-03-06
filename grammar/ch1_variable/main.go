package main

import (
	"fmt"
)

// 声明全局变量，全局变量可以不被使用
var g1, g2, g3 int
var g4 = 'a'

func main() {
	var a string = "Hello, world!"
	fmt.Println(a)

	// 同时声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 声明变量时不初始化，默认是零值
	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	// 零值是 nil
	var f *int
	fmt.Println(f)

	// 格式化输出
	fmt.Printf("%v %v %q\n", b, c, a)

	// 根据值自动判断类型
	var g = 1
	fmt.Println(g)

	// := 声明变量, 如果 var 已经声明过了，就会产生编译错误
	h := false
	fmt.Println(h)

	// 交换两个变量的值，类型必须相同
	fmt.Println(b, c)
	b, c = c, b
	fmt.Println(b, c)

	// 空白表示符 _ 用于抛弃值
	_, b = 1, 3
}
