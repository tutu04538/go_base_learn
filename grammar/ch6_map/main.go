package main

import "fmt"

func update(map1 map[string]int, s string, i int) {
	map1[s] = i
}

func main() {
	/*
		map：键值对
		使用前需要 make
		不能对 map 的元素取地址
		map不是并发安全的，需要使用 sync.Map 或互斥锁来保证并发安全
	*/

	map1 := make(map[string]int)
	fmt.Printf("%T\n", map1)

	map2 := map[string]int{
		"a": 1,
		"b": 2,
	}
	fmt.Printf("%v\n", map2)

	// 先声明，再 make
	var map3 map[string]int
	map3 = make(map[string]int)
	fmt.Printf("%T\n", map3)

	// map 使用
	map3["a"] = 1
	map3["b"] = 2
	map3["c"] = 3
	v1 := map3["a"]
	fmt.Println(v1)

	delete(map3, "a")

	for key, value := range map3 {
		fmt.Println(key, value)
	}

	// 判断键是否存在，不存在会返回零值
	v2, exists := map3["d"]
	if exists {
		fmt.Println("v2:", v2)
	} else {
		fmt.Println("v2 is nil: ", v2)
	}

	// 嵌套 map
	map4 := map[string]map[string]int{
		"a": {
			"a1": 1,
			"a2": 2,
		},
		"b": {
			"b1": 1,
			"b2": 2,
		},
	}
	fmt.Println(map4)

	// map 作为函数参数
	// map 是引用类型，作为函数参数传递时是指针
	update(map3, "a", 114514)
	fmt.Println(map3)

	// 初始化时指定容量
	map5 := make(map[string]int, 100)
	map5["a"] = 1

	// 随机遍历
	for i := 0; i < 3; i++ {
		for k, v := range map3 {
			fmt.Printf("%v:%v ", k, v)
		}
		fmt.Println()
	}

}
