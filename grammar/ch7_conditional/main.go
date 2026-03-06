package main

import (
	"fmt"
)

func main() {
	fmt.Println("conditional statement")

	age := 18
	if age >= 18 {
		fmt.Println("1111")
	} else {
		fmt.Println("2222")
	}

	if num := 9; num <= 10 {
		fmt.Println("3333")
	}

	// switch 语句自带 break 效果
	day := 3
	switch day {
	case 1:
		fmt.Println("day 1")
	case 2:
		fmt.Println("day 2")
	default:
		fmt.Println("not day 1 or day 2")
	}

	day = 6
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("workday")
	default:
		fmt.Println("weekend")
	}

	switch {
	case day <= 5:
		fmt.Println("workday")
	case day > 5:
		fmt.Println("weekend")
	}

	// 如果命中当前 case 后，需要继续执行下一个 case，可以使用 fallthrough
	switch {
	case day <= 5:
		fmt.Println("workday")
		fallthrough
	case day > 5:
		fmt.Println("weekend with fallthrough")

	}

	// 用 switch 进行类型判断

	var i interface{} = "Hello"
	switch v := i.(type) {
	case string:
		fmt.Printf("string: %s\n", v)
	case int:
		fmt.Printf("int: %d\n", v)
	}
}
