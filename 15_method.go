/*
 * @Author: Theo_hui
 * @Date: 2021-07-31 15:59:30
 * @Descripttion:
 */
package main

import "fmt"

//定义结构体
type rect struct {
	width, height int
}

// 结构体的方法一
// func (接收者,接收者类型) 方法名称（参数列表）（返回值列表）
func (r *rect) area() int {
	if r.height == 0 || r.width == 0 {
		return 0
	} else {
		return r.height * r.height
	}
}

// 方法二
func (r *rect) perim() int {
	if r.height == 0 || r.width == 0 {
		return 0
	} else {
		return 2 * (r.height + r.height)
	}
}

func main() {
	// 创建一个实例
	r := rect{width: 10, height: 20}

	//直接使用实例的方法
	fmt.Println(r.area())
	fmt.Println(r.perim())

	//指针的方式
	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.area())
}
