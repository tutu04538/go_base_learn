package main

import "fmt"

type ID interface {
	SetID(id int)
	GetID() int
}

type People struct {
	id int
}

func (p *People) SetID(id int) {
	p.id = id
}

func (p *People) GetID() int {
	return p.id
}

func GetID(r ID) {
	fmt.Printf("ID is %d\n", r.GetID())
}

type Name interface {
	ID
	GetName() string
}

type worker struct {
	People
	name string
}

func (w *worker) GetName() string {
	return w.name
}

func main() {

	p := People{
		id: 1,
	}
	fmt.Println(p.GetID())

	var id ID
	// id = new(People)
	id = &People{}
	id.(*People).id = 2
	fmt.Println(id.GetID())

	// 空接口
	// 任意类型都实现了空接口，因此空接口可以存储任意类型的数值
	var any interface{}
	any = 10
	fmt.Println(any)

	any = "tutu"
	fmt.Println(any)

	// 类型断言，用来检查接口变量的值是否实现了某个接口或者是否是某个具体的类型
	// value, ok := x.(T)
	// T 是某个具体的类型，类型断言会判断 x 的类型是否是 T
	// 下面的代码会报错
	/*
		var a int = 1
		var b interface{} = a
		var c int = b
	*/
	var x interface{}
	x = 8
	val, ok := x.(int)
	fmt.Printf("val is %d, ok is %t\n", val, ok)

	// interface 作为函数参数
	p2 := &People{
		id: 114514,
	}
	GetID(p2)

	// interface 嵌套
	w := &worker{
		name:   "tutu114515",
		People: People{id: 114515},
	}

	var wi Name = w
	fmt.Println(wi.GetName())
}
