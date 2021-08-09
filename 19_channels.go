/*
 * @Author: Theo_hui
 * @Date: 2021-08-09 14:41:56
 * @Descripttion:
 */
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true
}

// 设定通道方向 只写
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 设定方向，只读ping 只写ping
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	/*************无缓冲通道***********/
	//1. 创建通道
	message := make(chan string)

	//2. 匿名函数创建协程，发送消息
	go func() { message <- "ping" }()

	//3. main协程 接收消息
	msg := <-message
	fmt.Println(msg)

	/************有缓冲通道************/
	message = make(chan string, 2)

	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)

	/************通道用于协程同步***********/
	done := make(chan bool, 1)
	go worker(done)

	<-done

	/************通道方向*****************/
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs) //数据流 msg-[pings]->ping-[pongs]->pong
	/************通道选择器**************/
	c1 := make(chan string)
	c2 := make(chan string)

	//协程1
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	//协程2
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
