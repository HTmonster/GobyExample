/*
 * @Author: Theo_hui
 * @Date: 2021-07-31 10:39:31
 * @Descripttion:
 */
package main

import "fmt"

// 定义结构体
type person struct {
	name string
	age  int
}

func main() {

	// 创建结构体实例
	// 1
	xiaoming := person{"xiaoming", 11}
	fmt.Println(xiaoming)

	//2
	xiaohong := person{name: "xiaohong", age: 12}
	fmt.Println(xiaohong)
	fmt.Println(xiaohong.name)
}
