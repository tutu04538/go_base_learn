package main

import (
	"errors"
	"fmt"
)

// type error interface {
// 	Error() string
// }

func AddForPositiveNum(a int) (int, error) {
	if a > 0 {
		return a + 1, nil
	} else {
		return 0, fmt.Errorf("num is not a positive number")
	}
}

// 自定义 error 对象
type MyError struct {
	code int
	msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error msg: %s, error code: %d", e.msg, e.code)
}

func NewError(msg string, code int) error {
	return &MyError{msg: msg, code: code}
}

func GetCode(err error) int {
	if e, ok := err.(*MyError); ok {
		return e.code
	}
	return -1
}

func main() {
	err1 := errors.New("error_test")
	fmt.Println(err1)
	fmt.Printf("err1 type: %T\n", err1)

	err4 := errors.New("error_test")
	fmt.Println(err1.Error() == err4.Error()) // true

	b, err2 := AddForPositiveNum(1)
	if err2 != nil {
		fmt.Printf("Error: %s\n", err2)
	} else {
		fmt.Println(b)
	}

	c, err3 := AddForPositiveNum(-1)
	if err3 != nil {
		fmt.Printf("Error: %s\n", err3)
	} else {
		fmt.Println(c)
	}

	// 自定义 error 对象
	err5 := NewError("new error struct", 1)
	fmt.Println(err5)
	fmt.Println(GetCode(err5))

}
