package main

import "fmt"

// 定义结构体
type Student struct {
	ID    int
	Name  string
	Age   int
	Score int
}

func main() {

	// 结构体初始化
	st := Student{
		ID:   100,
		Name: "tutu",
	}
	fmt.Println(st)
	fmt.Printf("st1: %v\n", st)

	st2 := Student{
		100,
		"tutu2",
		1,
		1,
	}
	fmt.Printf("st2: %v\n", st2)

	// 访问结构体成员
	name := st.Name
	fmt.Println(name)
}
