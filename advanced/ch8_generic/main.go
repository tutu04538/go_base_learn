package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// 类型约束
func max[T interface{ int | float64 }](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

type Node[T any] struct {
	node  *Node[T]
	value T
}

func (t *Node[T]) addNode(v T) {
	n := &Node[T]{
		value: v,
	}
	t.node = n
}

// interface 不仅可以定义 method set，还可以定义 type set

type Value interface {
	int | float32
}

// MultiplyEach 返回切片中每个元素都乘以factor的副本切片
func MultiplyEachA[E constraints.Integer](s []E, factor E) []E {
	result := make([]E, len(s))
	for i, v := range s {
		result[i] = v * factor
	}
	return result
}

// ~T 运算符用于匹配所有以 T 为底层类型的类型集合
// 例如 type MyInt int, MyInt 的底层类型是 int
// 对于下面的例子，自定义一个 vector 类型，底层类型是 []int32
// type vector []int32
// 那么 S 就是 vector，最后返回的结果也是 vector，而不是一个普通的 []int32，因此定义在 vector 类型上的方法也都可以在返回值上继续使用
// 约束类型推断，S 依赖于 E，当 S 类型确定时，E 类型也就确定了，例如传入 []int32，那么 E 就是 int32
func MultiplyEachB[S ~[]E, E constraints.Integer](s S, factor E) S {
	result := make(S, len(s))
	for i, v := range s {
		result[i] = v * factor
	}
	return result
}

func main() {
	fmt.Println(max(1, 2))
	fmt.Println(max(1.2, 2.2))

	// 函数参数类型推断
	// 下面的两种写法等价
	max[int](1, 2)
	max(1, 2)
}
