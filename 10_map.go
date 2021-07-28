/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 18:06:36
 * @Descripttion:
 */
package main

import "fmt"

func main() {
	//创建
	m := make(map[string]int, 5) //大小可选，最好指定，避免资源浪费

	//赋值
	m["k1"] = 1
	m["k2"] = 10
	fmt.Println(m)

	//取值
	v1 := m["k1"]
	fmt.Println(v1)
	v2, exist := m["k2"]
	if exist {
		fmt.Println(v2)
	}

	//删除
	delete(m, "k2")
	fmt.Println(m)

	//取长度
	fmt.Println(len(m))

	//创建带初始化
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
