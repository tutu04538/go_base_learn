package main

import "fmt"

type Student struct {
	ID   int
	age  int
	name string
}

func (st Student) GetName() string {
	return st.name
}

func (st *Student) SetName(name string) {
	st.name = name
}

type People struct {
	age  int
	name string
}

type worker struct {
	ID int
	People
}

func (p *People) GetName() string {
	return p.name
}

func (w *worker) SetID(id int) {
	w.ID = id
}

func (w *worker) GetID() int {
	return w.ID
}

func main() {
	// 定义在值类型或指针类型上的方法，可以由值类型或指针类型的变量来调用
	st := &Student{
		ID:   100,
		age:  1,
		name: "tutu",
	}

	st.SetName("tutu2")

	fmt.Println(st.GetName())

	// 组合
	worker1 := worker{
		ID: 1,
		People: People{
			age:  33,
			name: "tutu3",
		},
	}

	fmt.Println(worker1.GetName())
	worker1.SetID(10)
	fmt.Println(worker1.ID)
}
