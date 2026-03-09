package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name  string
	age   int
	score float64
}

func (s *student) SetAge(a int) {
	s.age = a
}

func (s *student) GetAge() int {
	return s.age
}

func main() {

	a := 100
	t1 := reflect.TypeOf(a)
	fmt.Println(t1.String())
	v1 := reflect.ValueOf(a)
	fmt.Println(v1)
	fmt.Println(v1.String())

	b := student{
		name:  "tutu",
		age:   1,
		score: 90.5,
	}
	t2 := reflect.TypeOf(b)
	fmt.Println(t2.String())
	v2 := reflect.ValueOf(b)
	fmt.Println(v2)
	fmt.Println(v2.String())

	fmt.Printf("the field num of student is %d\n", v2.NumField())
	fmt.Printf("field1 type is %v, value is %s\n", v2.Field(0).Type().Name(), v2.Field(0).String())
	fmt.Printf("field3 type is %v, value is %.2f\n", v2.Field(2).Type().Name(), v2.Field(2).Float())

	// 通过反射调用方法

	c := &b
	t3 := reflect.TypeOf(c)
	v3 := reflect.ValueOf(c)
	for i := range t3.NumMethod() {
		fmt.Printf("%v\n", t3.Method(i))
	}

	m, _ := t3.MethodByName("SetAge")

	args := make([]reflect.Value, 0)
	args = append(args, v3)
	args = append(args, reflect.ValueOf(114514))
	m.Func.Call(args)
	fmt.Println(b)

	// CanAddr CanSet 没搞明白
}
