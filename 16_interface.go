/*
 * @Author: Theo_hui
 * @Date: 2021-07-31 17:34:44
 * @Descripttion:
 */
package main

import (
	"fmt"
	"math"
)

//定义接口 几何
// 几何为有面积有周长的东西
type geometry interface {
	area() float64
	perim() float64
}

//定义结构体
// 矩阵
type react struct {
	width, height float64
}

func (r react) area() float64 {
	return r.height * r.width
}
func (r react) perim() float64 {
	return 2 * (r.height + r.width)
}

// 圆形
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := react{20.0, 30.1}
	c := circle{11.2}

	measure(r)
	measure(c)

}
