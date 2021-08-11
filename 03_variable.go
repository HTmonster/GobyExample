/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 10:16:35
 * @Descripttion:
 */
package main

import "fmt"

func main() {
	// 创建变量
	var var_1 int               //默认赋值
	var var_2 = "hello"         //忽略类型直接赋值
	var var_3_1, var_3_2 = 1, 2 //分别赋值
	var_4 := 3.14159            //海象运算符 简写 声明并赋值

	fmt.Println(var_1)            //默认赋值为0
	fmt.Println(var_2)            //
	fmt.Println(var_3_1, var_3_2) //
	fmt.Println(var_4)            //
}
