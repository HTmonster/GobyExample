/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 15:39:39
 * @Descripttion:
 */
package main

import "fmt"

func main() {
	//1. 创建空的数组
	var a [5]int
	fmt.Println("array a:", a) //int 零值

	//修改值
	a[4] = 100
	fmt.Println("array a:", a)

	//2. 创建带初始化的数组
	b := [5]int{2, 4, 1, 5, 3}
	fmt.Println("array b", b)

	//3. 二维数组
	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("array c:", c)
}
