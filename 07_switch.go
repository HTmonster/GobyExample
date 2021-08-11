/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 14:48:26
 * @Descripttion:
 */
package main

import (
	"fmt"
	"time"
)

func main() {

	// 1. 普通用法 判断值类型
	name := "hong"
	fmt.Printf("Welcome:")
	switch name {
	case "min":
		fmt.Println("MIN")
	case "hong":
		fmt.Println("HONG")
	default:
		fmt.Println("....")
	}

	// 2. 值可以进行合并
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is weekday")
	}

	// 3. 可以为表达式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before 12")
	default:
		fmt.Println("It's after 12")
	}

	// 4. 可以为类型
	WhatAmI := func(i interface{}) {

		switch i.(type) {
		case nil:
			fmt.Println("I'm nil")
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		case float64:
			fmt.Println("I'm float64")
		case string:
			fmt.Println("I'm string")
		default:
			fmt.Println("sorry")
		}
	}
	WhatAmI(1)
	WhatAmI("string")
	WhatAmI(3.14)

	// fallthrough
	score := 80
	fmt.Printf("您的奖品有：")
	switch {
	case score >= 60:
		fmt.Printf(" 笔*2")
		fallthrough
	case score >= 80:
		fmt.Printf(" xbox游戏机*1")
		fallthrough
	default:
		fmt.Printf(" 迪士尼豪华游*1")
	}

}
