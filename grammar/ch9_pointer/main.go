package main

import "fmt"

func main() {

	var a *int
	var b int = 3
	a = &b
	fmt.Printf("%p\n", a)

	// new(type)
	s := new(string)
	fmt.Printf("%p\n", s)

	*s = "abc"
	fmt.Printf("%v\n", *s)
}
