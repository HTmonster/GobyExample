/*
 * @Author: Theo_hui
 * @Date: 2021-08-11 10:16:52
 * @Descripttion:
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	/****************Timer 定时器*****************/
	// 1. After函数
	go func(ch <-chan time.Time) {
		// 打印接收到的信息
		fmt.Printf("Now is %s\n", <-ch)
		done <- true
	}(time.After(2 * time.Second)) //设置2秒的定时器，然后将返回值发送给协程

	<-done
	// 2. AfterFunc函数
	go time.AfterFunc(time.Second*2, func() {
		fmt.Println("func invoked after 2 seconds")
		done <- true
	})
	<-done
	//3. NewTimer函数
	timer := time.NewTimer(2 * time.Second) //生成一个定时器

	//timer.C 返回通道
	<-timer.C
	fmt.Println("2 seconds later...")

	/**************Ticker 打点器***************/
	// var wg sync.WaitGroup

	// // 1.Tick函数
	// wg.Add(10)
	// go func(ch <-chan time.Time) {
	// 	for t := range ch {
	// 		fmt.Printf("Backup at: %s\n", t)
	// 		wg.Done()
	// 	}
	// }(time.Tick(time.Second * 1))
	// wg.Wait()

	// // 2.NewTicker函数
	ticker := time.NewTicker(500 * time.Millisecond)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true

}
