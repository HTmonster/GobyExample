/*
 * @Author: Theo_hui
 * @Date: 2021-08-20 17:21:36
 * @Descripttion:
 */
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//=======================编码========================
	//bool值
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	//整数值
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	//浮点数值
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	//字符串值
	strB, _ := json.Marshal("hahahah")
	fmt.Println(string(strB))

	//切片值
	slcB, _ := json.Marshal([]string{"apple", "peach"})
	fmt.Println(string(slcB))

	//map值
	mapB, _ := json.Marshal(map[string]int{"apple": 1, "peach": 2})
	fmt.Println(string(mapB))

	//结构体值
	type mystruct struct {
		Name string //只有首字母大写才能被编码
		Sex  string `json:"Gender"` //使用声明标签来定义json数据的键名
		age  int    //首字母不大写 不可导出
	}
	resB, _ := json.Marshal(&mystruct{"ming", "male", 20})
	fmt.Println(string(resB))

	//=====================解码================================
	//json数据
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//接收的数据结构
	var dat map[string]interface{}

	//解码
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	//类型转换
	num := dat["num"].(float64)
	fmt.Println(num)

	//访问嵌套值需要一系列的转换
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//2. 直接解释为自定义的数据结构
	str := `{"Name": "Hong", "Gender": "female"}`
	res := mystruct{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
