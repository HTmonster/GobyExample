/*
 * @Author: Theo_hui
 * @Date: 2021-07-31 09:46:22
 * @Descripttion:
 */
package main

import "fmt"

func main() {

	// 创建切片
	nums := []int{1, 3, 55}

	// range 对nums进行遍历
	for i, num := range nums {
		fmt.Println(i, num)
	}

	// 创建map
	kvs := map[string]int{"joke": 1, "lisa": 2}
	// range 对 kvs进行遍历
	for k, v := range kvs {
		fmt.Println(k, v)
	}
	// 只遍历键
	for k := range kvs {
		fmt.Println(k)
	}

	// 遍历字符串
	for i, c := range "hello" {
		fmt.Println(i, c)
	}

}
