/*
 * @Author: Theo_hui
 * @Date: 2021-08-20 17:06:26
 * @Descripttion:
 */
package main

import (
	"fmt"
	"regexp"
)

func main() {
	//1. 直接使用，不编译
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//2. 编译后使用
	r, _ := regexp.Compile("p([a-z]+)ch")
	//2.1 匹配字符串
	fmt.Println(r.MatchString("peach"))
	//2.2 查找字符串
	fmt.Println(r.FindString("a peach"))
	fmt.Println(r.FindStringIndex("a peach"))
	fmt.Println(r.FindStringSubmatch("a peach"))
	fmt.Println(r.FindStringSubmatchIndex("a peach"))
	fmt.Println(r.FindAllString("a peach", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("a peach", -1))
	//2.3 字符串替换
	fmt.Println(r.ReplaceAllString("a peach", "<friut>"))
}
