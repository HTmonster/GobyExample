/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 10:03:49
 * @Descripttion:
 */

package main

import "fmt"

func main() {
	// 1. 布尔类型
	fmt.Println(true)
	fmt.Println(true && true)
	fmt.Println(true || false)

	// 2. 数字类型
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	// 3. 字符类型
	fmt.Println('A')
	fmt.Println('\u0056')

	// 4. 字符串类型
	fmt.Println("go" + "lang")
}
