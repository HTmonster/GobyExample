/*
 * @Author: Theo_hui
 * @Date: 2021-08-20 17:49:39
 * @Descripttion:
 */
package main

import (
	"encoding/xml"
	"fmt"
)

// 定义结构体
type Plant struct {
	XMLName xml.Name `xml:"Plant"`   //xml元素名称
	Id      int      `xml:"id,attr"` //id是一个属性，而不是嵌套元素
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func main() {
	//创建一个结构体
	coffee := &Plant{Id: 27, Name: "coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	//进行编码
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	//添加头部
	fmt.Println(xml.Header + string(out))

	//解码
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

}
