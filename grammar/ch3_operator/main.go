package main

import "fmt"

func main() {

	// 算术运算符
	var a, b = 1, 2
	var c int

	c = a + b
	c = a - b
	c = a * b
	c = a / b
	fmt.Printf("%d\n", c) // output: 0
	c = a % b
	fmt.Printf("%d\n", c) // output: 1
	a++
	fmt.Printf("%d\n", a)
	a = 21
	a--
	fmt.Printf("%d\n", a)

	// 关系运算符
	// ==
	// !=
	// >
	// <
	// >=
	// <=

	// 逻辑运算符
	// && || !
	var d, e = true, false
	if d && e {
		fmt.Println("true1")
	}

	if d || e {
		fmt.Println("true2")
	}

	// 位运算符
	// & | ^ << >>
	var f uint = 60
	var g uint = 2
	var h uint

	h = f & g
	h = f | g
	h = f ^ g
	h = f << g
	h = f >> g
	fmt.Println(h)

	// 赋值运算符
	var i uint = 3
	var j uint = 4
	var k uint

	k = i
	k *= i
	k <<= j
	fmt.Println(k)

	// 取地址运算符 & 和 指针运算符 *

	var l uint = 3
	var m float64 = 0.2
	fmt.Printf("type of var l: %T\n", l)
	fmt.Printf("type of var m: %T\n", m)

	var ptr *uint = &l
	fmt.Printf("type of var ptr: %T\n", ptr)
	fmt.Printf("the value of *ptr: %d\n", *ptr)

	// 运算符优先级
	// 1: * / % << >> & | ^
	// 2: + -
	// 3: == != < <= > >=
	// 4: &&
	// 5: ||
}
