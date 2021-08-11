/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 10:36:53
 * @Descripttion:
 */
package main

import (
	"fmt"
	"math"
)

//字符串常量
const s string = "constant"

func main() {

	fmt.Println(s)

	// 常数表达式
	const n = 500000000
	const d = 3e20 / n

	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n)) //自动转为float64类型
}
