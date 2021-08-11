/*
 * @Author: Theo_hui
 * @Date: 2021-08-11 13:57:00
 * @Descripttion:
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

// 工作协程 id 只读通道jobs 只写通道 results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println(">> worker[", id, "] start job", j)
		time.Sleep(1 * time.Second)
		fmt.Println("   << worker[", id, "] finished job", j)
		results <- j
	}
}

// 工作协程 waitgroup实现
func worker2(id int, wg *sync.WaitGroup) {

	defer wg.Done() //不管如何都要保证完成

	fmt.Printf("[%d] worker ->\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("\t<- worker [%d]\n", id)
}

func main() {
	/*************************普通实现方法******************/
	const jobnums = 5
	jobs := make(chan int, jobnums)
	results := make(chan int, jobnums)

	// 创建三个工人
	for w := 1; w < 3; w++ {
		go worker(w, jobs, results)
	}

	// 生成任务
	for i := 1; i <= jobnums; i++ {
		jobs <- i
	}
	close(jobs)
	// 等待生成结果
	for j := 1; j <= jobnums; j++ {
		<-results
	}

	/************************waitgroup方法****************/
	var wg sync.WaitGroup //创建组

	// 分配工人工作
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker2(i, &wg)
	}

	// 等待所有完成
	wg.Wait()
}
