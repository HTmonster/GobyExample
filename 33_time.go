/*
 * @Author: Theo_hui
 * @Date: 2021-08-23 09:53:39
 * @Descripttion:
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	/*========time包的常见操作========*/
	now := time.Now()
	fmt.Println(now) //当前时间
	then := time.Date(2029, 10, 7, 20, 34, 58, 12345456789, time.Now().Local().Location())
	fmt.Println(then)

	/*=======time数据结构============*/
	// 组成部分
	fmt.Println(then.Year(), then.Location())
	// 时间前后比较
	fmt.Println(then.After(now))
	fmt.Println(then.Before(now))
	fmt.Println(then.Equal(now))
	//时间段
	diff := then.Sub(now)
	fmt.Println(diff)
	/*-------Duration数据结构----------*/
	fmt.Println(diff.Hours(), diff.Microseconds())
	// 时间加上时间段
	fmt.Println(now.Add(diff))

	/*=======时间戳=================*/
	fmt.Println(now.Unix(), now.UnixNano())
	fmt.Println(time.Unix(100000000, 0))

	/*=======格式化=================*/
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Format("3:04PM"))
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	fmt.Println(t1, e)
	fmt.Println(time.Parse("3 04 PM", "8 41 PM"))
}
