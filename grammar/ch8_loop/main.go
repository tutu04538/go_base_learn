package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		for init; condition; post {
		}

		for condition {
		}

		for {
		}
	*/

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// for range
	// 可以对 array slice map string 进行迭代循环
	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	for key, value := range map1 {
		fmt.Printf("%s:%d\n", key, value)
	}

	// for range 的坑
	// 取元素的地址，下面的代码取不到原元素的地址
	// for range 会创建每个元素的副本，而不是直接操作原始切片中的元素
	slice1 := []int{1, 2, 3}
	ptrs := []*int{}
	for _, v := range slice1 {
		fmt.Printf("%p\n", &v)
		ptrs = append(ptrs, &v)
	}

	for _, v := range ptrs {
		fmt.Printf("%d\n", *v)
	}

	a := ptrs[0]
	*a = 114514
	fmt.Printf("本来的地址：%p\n", &slice1[0])
	fmt.Printf("for循环取到的地址：%p\n", a)
	fmt.Println(slice1[0])
	fmt.Printf("for循环取到的地址不是本来的地址\n")

	// 循环是否会停止
	slice2 := []int{1, 2, 3}
	for _, v := range slice2 {
		slice2 = append(slice2, v)
	}
	fmt.Println(len(slice2))
	fmt.Println("循环会停止")

	// 使用迭代变量时的闭包问题
	var funcs []func()
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	for _, f := range funcs {
		f()
	}

	// 遍历字典时的顺序，使用sort对键进行排序后，再遍历
	map2 := map[string]int{"a": 1, "b": 2, "c": 3}

	keys := []string{}
	for k := range map2 {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s:%d ", k, map2[k])
	}
	fmt.Println()

	// 遍历字符串
	s1 := "hello 世界"
	for i, str := range s1 {
		fmt.Printf("idx %d: %c\n", i, str)
	}
	/*
		for range 遍历的是 Unicode 代码点，而不是字节
		输出：
		idx 0: h
		idx 1: e
		idx 2: l
		idx 3: l
		idx 4: o
		idx 5:
		idx 6: 世
		idx 9: 界
	*/
	fmt.Println()
	for i := 0; i < len(s1); i++ {
		fmt.Printf("idx %d: %x\n", i, s1[i])
	}
	/*
		idx 0: 68
		idx 1: 65
		idx 2: 6c
		idx 3: 6c
		idx 4: 6f
		idx 5: 20
		idx 6: e4
		idx 7: b8
		idx 8: 96
		idx 9: e7
		idx 10: 95
		idx 11: 8c
	*/
}
