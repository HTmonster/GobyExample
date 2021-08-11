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

	/***************超时处理******************/
	// 通道一与协程一
	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second) //延时2秒
		ch1 <- "func 1"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout:1 second") //先被触发
	}

	//通道二与协程二
	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second) //延时1秒
		ch2 <- "func 2"
	}()

	select {
	case res := <-ch2:
		fmt.Println(res) //先被触发
	case <-time.After(2 * time.Second):
		fmt.Println("timeout:2 second")
	}

	/*****************非阻塞通道*********************/
	message2 := make(chan string) //无缓冲
	signal2 := make(chan bool)    //无缓冲

	// 1. message2 通道没有消息, 非阻塞接收
	select {
	case msg := <-message2:
		fmt.Println(msg)
	default:
		fmt.Println("no message received")
	}

	// 2. message2 通道有消息，非阻塞发送
	msg2 := "hello"
	select {
	case message2 <- msg2:
		fmt.Println(msg2)
	default:
		fmt.Println("no message sent")
	}

	// 3. message2 和 signal2 非阻塞接收
	select {
	case msg := <-message2:
		fmt.Println(msg)
	case sig := <-signal2:
		fmt.Println(sig)
	default:
		fmt.Println("no activity")
	}

	/****************** 通道地图 **************/
	jobs := make(chan int, 5)
	done3 := make(chan bool)

	//协程 处理任务
	go func() {
		for {
			j, more := <-jobs
			if more {
				// 还有任务
				fmt.Println("working...", j)
			} else {
				// 没有任务
				fmt.Println("no more work")
				done3 <- true
				return
			}
		}
	}()

	//主协程 发送任务
	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	// 关闭通道
	close(jobs)
	fmt.Println("sent all jobs")

	<-done3
	/******************通道遍历*****************/
	//创建缓存通道 并发送消息
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//通道读取
	for item := range queue {
		fmt.Println(item)
	}
}
