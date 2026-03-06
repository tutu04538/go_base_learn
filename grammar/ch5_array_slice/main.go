package main

import (
	"fmt"
)

func main() {
	/*
		数组特点
		1、长度固定
		2、元素类型一致
		3、长度也是类型，例如 [3]int 和 [5]int 是不同的类型
	*/
	// 数组定义
	var scores [3]int
	fmt.Println(scores)
	var prices = [3]float64{1.1, 1.2, 1.3}
	fmt.Printf("%v\n", prices)
	names := [...]string{"a", "b", "c"}
	fmt.Println(names)
	colors := [5]int{0: 1, 2: 114514}
	fmt.Println(colors) // [1 0 114514 0 0]

	sum := 0.0
	for _, price := range prices {
		sum += price
	}
	average := sum / float64(len(prices))
	fmt.Printf("average: %.2f\n", average) // 1.20

	/*
		切片：大小可以动态调整
		三个核心概念
		1、指针：指向底层数组的第一个元素
		2、长度：切片当前的元素个数 len
		3、容量：从切片起始位置到底层数字末尾的元素个数 cap
		注意事项
		1、切片是引用类型，多个切片可能共享同一个底层数组
		2、append可能导致重新分配内存，生成新的底层数组
		3、使用make创建切片时，可以指定容量来减少内存重新分配的次数
	*/

	// 直接创建
	slice1 := []string{"a", "b", "c"}
	fmt.Println(slice1)

	// 通过 make 创建，make 用于预先分配内存
	slice2 := make([]int, 3, 5) // 长度为3，容量为5的切片
	fmt.Println(slice2)

	// 从数组中创建
	arr := [5]int{1, 2, 3, 4, 5}
	slice3 := arr[0:4] // 长度为4的切片
	fmt.Println(slice3)
	fmt.Printf("type of arr: %T\n", arr)       // [5]int
	fmt.Printf("type of slice3: %T\n", slice3) // []int

	// 实际使用
	slice4 := []string{"a"}
	slice4 = append(slice4, "b", "c")
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice5 := make([]string, len(slice4))
	copy(slice5, slice4)
	fmt.Println(slice5)

	slice6 := slice4[1:2]
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))
	fmt.Println(slice4)
	slice6[0] = "test1"
	fmt.Println(slice6)
	fmt.Println(slice4)
	slice6 = append(slice6, "test2")
	fmt.Println(slice4)
	fmt.Println("打印地址，二者指向同一块内存")
	fmt.Printf("%p\n", slice4)
	fmt.Printf("%p\n", slice6)

	fmt.Println("此时 slice6 的 len 和 cap 相等，继续 append")
	slice6 = append(slice6, "test3")
	fmt.Println(slice4)
	fmt.Println(slice6)
	fmt.Println(len(slice6))
	fmt.Println(cap(slice6))
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println("打印地址，二者不再指向同一块内存")
	fmt.Printf("%p\n", slice4)
	fmt.Printf("%p\n", slice6)
}
