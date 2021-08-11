/*
 * @Author: Theo_hui
 * @Date: 2021-08-11 14:54:23
 * @Descripttion:
 */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	/*******************原子计数*****************/
	//创建计数变量
	var ops uint64

	//WaitGroup同步协程
	var wg sync.WaitGroup

	//创建50个协程
	for i := 1; i <= 50; i++ {
		wg.Add(1) //增加监控一个协程
		go func() {
			// 每个协程 计数1000次
			for j := 0; j < 1000; j++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)

}
