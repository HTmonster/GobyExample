/*
 * @Author: Theo_hui
 * @Date: 2021-07-31 10:32:27
 * @Descripttion:
 */
package main

import "fmt"

/* 两种传递值的方法 */

// 1. 传递值
func zeroval(val int) {
	val = 0 //尝试修改
}

// 2. 传递指针
func zeroptr(val *int) {
	*val = 0 //尝试修改
}

func main() {
	i := 1

	//第一种方法
	zeroval(i)
	fmt.Println(i)

	//第二种方法
	zeroptr(&i)
	fmt.Println(i)

}
