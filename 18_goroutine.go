/*
 * @Author: Theo_hui
 * @Date: 2021-08-09 14:14:38
 * @Descripttion:
 */
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// 默认 main goroutine

	// 1.1 直接调用函数(阻塞式：会等待返回)
	f("direct")

	// 1.2 使用协程调用函数 (非阻塞式：不会等待返回)
	go f("goroutine")

	// 2 使用匿名函数创建协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) //最好使用WaitGroup
	fmt.Println("done")
}
