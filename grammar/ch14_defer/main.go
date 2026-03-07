package main

import (
	"fmt"
	"io"
	"os"
)

func defer1() {
	fmt.Println("defer1")
}

func defer2() {
	fmt.Println("defer2")
}

func defer3() {
	fmt.Println("defer3")
}

// 使用 defer 释放资源
func CopyFile(dst string, src string) (int64, error) {
	srcFd, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	// 问题：如果 Create 失败并直接返回，则 srcFd 没有被清理
	// 解决方法：使用 defer 延迟 srcFd 的清理
	defer srcFd.Close()

	dstFd, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFd.Close()

	num, err := io.Copy(dstFd, srcFd)
	return num, err
}

// 测试 defer 中的函数的参数是否会在之后被修改
func deferRun1(a int) {
	defer fmt.Printf("deferRun1: %d\n", a) // 传值
	a += 1
	// deferRun1: 1
}

func PrintArr(arr *[4]int) {
	for _, v := range arr {
		fmt.Printf("deferRun2: %d\n", v)
	}
}

func deferRun2() {
	arr := [4]int{1, 2, 3, 4}
	defer PrintArr(&arr) // 传地址
	arr[0] = 114514
	/*
		deferRun2: 114514
		deferRun2: 2
		deferRun2: 3
		deferRun2: 4
	*/
}

func addAfterReturn() (res int) {
	a := 1
	defer func() {
		res++
	}()
	return a
}

func deferRun3() {
	fmt.Printf("deferRun3: %d\n", addAfterReturn())
	/*
		deferRun3: 2
		return 机制：
		1、设置返回值
		2、执行 defer 语句
		3、 将结果返回
	*/
}

func main() {
	/*
		defer 用于延迟函数或方法的调用
		defer 只能用于函数或方法前面
		defer 延迟的函数会在 return 或 panic 时被调用
		defer 延迟的函数遵循先进后出 LIFO
		defer 定义的延迟函数的参数在 defer 语句出时就已将确定下来了
	*/

	defer defer1()
	defer defer2()
	defer defer3()
	// output:
	// defer3
	// defer2
	// defer1
	c := 1
	deferRun1(c)
	deferRun2()
	deferRun3()

	// 使用 defer 处理异常
	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	a := 1
	b := 0
	fmt.Println(a / b)
	// runtime error: integer divide by zero

}
