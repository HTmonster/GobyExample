/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 14:24:24
 * @Descripttion:
 */

package main

import "fmt"

func main() {

	//1. 基本的结构
	i := 1
	for i < 10 {
		fmt.Println("i->", i)
		i += 1
	}

	//2. 类C结构
	for j := 0; j < 10; j++ {
		fmt.Println("j->", j)
	}

	//3. range循环
	str1 := "hello"
	for i, c := range str1 {
		fmt.Printf("%d:%c\n", i, c)
	}

	//beak
	for k := 0; k < 100; k++ {
		fmt.Println("k->", k)
		break
	}

	//continue
	for l := 0; l < 10; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Println("l->", l)
	}
}
