/*
 * @Author: Theo_hui
 * @Date: 2021-08-06 08:40:20
 * @Descripttion:
 */
package main

import (
	"errors"
	"fmt"
)

//函数返回值返回错误
func test1(arg int) (int, error) {

	if arg < 0 {
		return -1, errors.New("arg must bigger than 0")
	}
	return arg, nil
}

// 结构体 自定义错误类型
type argError struct {
	arg  int
	prob string
}

// 通过方法实现Error方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// 返回自定义的错误
func test2(arg int) (int, error) {
	if arg < 0 {
		return -1, &argError{arg, "arg must bigger than 0"}
	}

	return arg, nil
}

func main() {

	for _, i := range []int{-10, 10} {
		if r, e := test1(i); e == nil {
			fmt.Println("test1 work ", r)
		} else {
			fmt.Println("test1 fail ", e)
		}
	}

	for _, i := range []int{-10, 10} {
		if r, e := test2(i); e == nil {
			fmt.Println("test1 work ", r)
		} else {
			fmt.Println("test1 fail ", e)
		}
	}

	// 通过类型断言来得到自定义错误类型的实例
	_, e := test2(-100)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
