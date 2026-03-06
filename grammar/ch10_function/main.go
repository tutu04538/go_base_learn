package main

import (
	"fmt"
	"math"
)

func swap(x int, y int) (int, int) {
	return y, x
}

// 定义一个 fc 类型
type fc func(int) int

func callback(x int) int {
	fmt.Printf("callback: %d\n", x)
	return x
}

func Callback(x int, f fc) {
	f(x)
}

func getNumber() func() int {
	i := 1
	return func() int {
		i += 1
		return i
	}
}

func main() {
	x, y := 1, 2
	a, b := swap(x, y)
	fmt.Println(a, b)

	// 函数变量
	getSqrt := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSqrt(1.21))

	// 函数变量做回调函数
	Callback(1, callback)

	// 函数闭包
	nextNumber := getNumber()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}
