/*
 * @Author: Theo_hui
 * @Date: 2021-08-11 15:12:53
 * @Descripttion:
 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	/*****************互斥锁*********************/
	// 临界区
	maps := make(map[int]int)
	// 互斥量
	mutex := &sync.Mutex{}

	// 记录读取数
	var readops, writeops uint64

	// 创建100个协程 来读
	for i := 1; i <= 100; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				mutex.Lock()                  //加锁
				fmt.Println(maps[key])        //读取
				mutex.Unlock()                //解锁
				atomic.AddUint64(&readops, 1) //记录读取数
				time.Sleep(time.Microsecond)  //暂停一下
			}
		}()
	}

	// 创建10个协程 来写
	for j := 1; j <= 10; j++ {
		go func() {
			key := rand.Intn(5)
			value := rand.Intn(100)
			mutex.Lock()                   //加锁
			maps[key] = value              //写入
			mutex.Unlock()                 //解锁
			atomic.AddUint64(&writeops, 1) //记录写入数
			time.Sleep(time.Microsecond)
		}()
	}

	// 主协程暂停一下
	time.Sleep(time.Second)
	// 显示读取和写入次数
	fmt.Println("writeops:", atomic.LoadUint64(&writeops))
	fmt.Println("readops:", atomic.LoadUint64(&readops))
	// 显示临界区
	mutex.Lock()
	fmt.Println(maps)
	mutex.Unlock()
}
