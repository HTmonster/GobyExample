/*
 * @Author: Theo_hui
 * @Date: 2021-08-11 14:28:40
 * @Descripttion:
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	/*******基本的资源限制********/
	// 1. 请求水桶 容量为5
	request := make(chan int, 5)
	// 1.1 模拟五个请求
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)

	// 2. 打点器 限制接收频率
	limiter := time.Tick(200 * time.Millisecond)
	// 3. 接收请求
	for req := range request {
		<-limiter
		fmt.Println("processing request ", req, time.Now())
	}

	/**********运行有短暂并发的限制************/
	// 1. 爆发事件通道 最多三个
	burstyLimiter := make(chan time.Time, 3)
	// 1.1 填充，表示允许的并发
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 1.2 每200ms尝试添加一个新值
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 2 模拟请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// 3 处理请求
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
