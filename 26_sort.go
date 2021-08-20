/*
 * @Author: Theo_hui
 * @Date: 2021-08-17 15:36:34
 * @Descripttion:
 */
package main

import (
	"fmt"
	"sort"
)

// 自定义的数据类型
type mylist []string

// 实现sort接口的三个方法
// 1. Len() 长度
func (l mylist) Len() int {
	return len(l)
}

// 2. Swap() 交换逻辑
func (l mylist) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// 3. Less() 判断逻辑
func (l mylist) Less(i, j int) bool {
	return len(l[i]) < len(l[j])
}

func main() {

	/*****************内置数据进行排序*************/
	strs := []string{"a", "hb", "jt", "ha"}
	ints := []int{1, 2, 3, 4, 5}

	// 对内置的数据进行排序
	sort.Ints(ints)
	fmt.Println(ints)
	sort.Strings(strs)
	fmt.Println(strs)

	//查看是否排序
	fmt.Println(sort.IntsAreSorted(ints))
	fmt.Println(sort.StringsAreSorted(strs))

	/****************自定义数据**********************/
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(mylist(fruits))
	fmt.Println(fruits)
}
