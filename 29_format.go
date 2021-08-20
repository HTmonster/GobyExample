/*
 * @Author: Theo_hui
 * @Date: 2021-08-20 14:47:54
 * @Descripttion:
 */
package main

import (
	"fmt"
)

type mystruct struct {
	name string
	age  int
}

func main() {
	//对象：各种数据类型
	i_struct := mystruct{"ming", 12}
	i_int := 1024
	i_float := 3.14159
	i_string := "hello"
	i_bool := true
	i_p := &i_struct

	//通用
	fmt.Println("============[通用]=================")
	fmt.Println("format\tstruct\t\t\t\t\tint\tfloat\tstring\tbool\tpointer")
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Printf("v:\t%v\t\t\t\t%v\t%v\t%v\t%v\t%v\n", i_struct, i_int, i_float, i_string, i_bool, i_p)
	fmt.Printf("+v:\t%+v\t\t\t%+v\t%+v\t%+v\t%+v\t%+v\n", i_struct, i_int, i_float, i_string, i_bool, i_p)
	fmt.Printf("#v:\t%#v\t%#v\t%#v\t%#v\t%#v\t%#v\n", i_struct, i_int, i_float, i_string, i_bool, i_p)
	fmt.Printf("T:\t%T\t\t\t\t%T\t%T\t%T\t%T\t%T\n", i_struct, i_int, i_float, i_string, i_bool, i_p)
	fmt.Println("-----------------------------------------------------------------------------------------")

	//整数
	fmt.Println("============[整数]=================")
	fmt.Printf("b\t\tc\td\to\tq\tx\tX\tU\n")
	fmt.Printf("%b\t%c\t%d\t%o\t%q\t%x\t%X\t%U\n", i_int, i_int, i_int, i_int, i_int, i_int, i_int, i_int)

	// 浮点数
	fmt.Println("============[浮点数]================")
	fmt.Printf("b\t\t\te\t\tE\t\tf\t\tF\t\tg\tG\t\n")
	fmt.Printf("%b\t%e\t%E\t%f\t%F\t%g\t%G\t\n", i_float, i_float, i_float, i_float, i_float, i_float, i_float)

	// 布尔值
	fmt.Println("============[布尔值]================")
	fmt.Printf("t:\t%t\n", i_bool)

	// 字符串
	fmt.Println("============[字符串]================")
	fmt.Printf("s\tq\tx\t\tX\n")
	fmt.Printf("%s\t%q\t%x\t%X\n", i_string, i_string, i_string, i_string)

	// 指针
	fmt.Println("=============[指针]=================")
	fmt.Printf("p:\t%p", i_p)
}
