package main

import (
	"fmt"
	"unsafe"
)

// const identifier [type] = value

// 用常量表示枚举类型
const (
	Unknown = 0
	Success = 1
	Fail    = 2
)

// 常量可以用 len、cap、unsafe.sizeof 函数来计算表达式的值，且必须是编译期可以确定的值
const (
	c1 = "abc"
	c2 = len(c1)
	c3 = unsafe.Sizeof(c1)
)

// iota 是一个特殊常量，在 const 关键字出现时将被重置为
// const 中每新增一行常量声明将使 iota 计数一次，可以将 iota 理解为 const 语句块中的行索引
const (
	i1 = iota
	i2 = iota
	i3 = iota
)

func main() {
	const b string = "abc"
	const c = "abc"
	const d, e = "abc", "abc"

	fmt.Println(c1, c2, c3)

	fmt.Println(i1, i2, i3)
	// 输出 0 1 2

	const (
		t1 = iota
		t2
		t3
		t4 = iota
		t5 = "abc"
		t6 // 和 t5 一致
		t7 = 100
		t8 // 和 t7 一致
		t9 = iota
	)

	fmt.Println(t1, t2, t3, t4, t5, t6, t7, t8, t9)
	// 0, 1, 2, 3, abc, abc, 100, 100, 8

	const (
		o1 = 1 << iota
		o2 = 3 << iota
		o3 // 实际上是 3 << 2
		o4 // 3 << 3
	)
	fmt.Println(o1, o2, o3, o4)
	// 1, 6, 12, 24

}
