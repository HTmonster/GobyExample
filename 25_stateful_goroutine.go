/*
 * @Author: Theo_hui
 * @Date: 2021-08-11 15:26:54
 * @Descripttion:
 */
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 读取操作
type readOp struct {
	key  int
	resp chan int //反馈
}

// 写入操作
type writeOp struct {
	key  int
	val  int
	resp chan bool //反馈
}

func main() {
	// 记录读取和写入次数
	var readOps, writeOps uint64

	// 读取和写入的通道
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 状态协程
	go func() {
		//独自享有 临界区
		state := make(map[int]int)
		// select来协同写入和读取
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 其它的协程 读取
	for i := 0; i < 100; i++ {
		go func() {
			for {
				//1.创建一个读操作
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				//2. 读操作加入通道
				reads <- read
				//3. 等待操作的结果返回
				<-read.resp
				// 记录读取数
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 其它的协程 写入
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//运行1s
	time.Sleep(time.Second)

	// 显示读取和写入次数
	fmt.Println("writeops:", atomic.LoadUint64(&writeOps))
	fmt.Println("readops:", atomic.LoadUint64(&readOps))
}
