#### :rocket: GobyExample

> 跟随 [Go by Example 中文版 (gobyexample-cn.github.io)](https://gobyexample-cn.github.io/) 项目复习Golang语言



#### :bicyclist: 完成清单

- 基本数据
	- [x] Hello world
	- [x] value
	- [x] variable 
	- [x] constant
- 基本控制
	- [x] for
	- [x] if/else
	- [x] switch
- 高级数据
	- [x] array
	- [x] slice
	- [x] map

#### :warning: 注意点

| 知识点 | 注意内容 | 链接 |
| ------ | -------- | ---- |
|   Golang值基本类型     |          |  [基本类型](#value)|
|   Golang变量     |          |  [变量](#variable)   |
|      Golang常量  |          | [常量](#const)     |
| Golang for循环||[for循环](#for)|
| Golang if/else判断|1. 花括号必需；2. 没有三目运算|[if/else判断](#ifelse)|
| Golang switch分支|1.任何类型 2.可以为表达式 3.break/fallthrough 4. 类型判断算法 |[switch分支](#switch)|
| Golang 数组|1.长度固定 2.空零值|[数组](#array)|
| Golang 切片|2.长度可以不固定 2. make初始化 3. append操作 4.copy操作|[切片](#slice)|
| Golang map||[map](#map)|

****


##### <span id="value">基本类型</span>

###### 布尔类型
- var b bool = true
- 逻辑运算符
	- 与 &&
	- 或 ||
	- 非 ！
###### 数字类型
- 基于架构的类型
	- int uint
		- 64位  64bit
		- 32位 32bit
	- uintptr
		- 足够存放一个指针即可
- 与操作系统架构无关类型
	- 有符号整数
		- int8
		- int16
		- int32
		- int64
	- 无符号整数
		- uint8
		- uint16
		- uint32
		- uint 64
	- 浮点数
		- float32
		- float64
- 进制
	- 077 八进制
	- 0xFF 十六进制
- 复数
	- complex64
	- complex128
- 位运算
	- 要求
		- 整数
		- 等位长
	- 二元运算符
		- 按位与 &
		- 按位或 ！
		- 按位异或 ^
	- 一元运算符
		- 按位补足 ^
		- 位左移 <<
		- 位右移 >>
###### 字符类型
- ASCII字符
	- byte
- Unicode字符
	- int \u十六进制
###### 字符串类型
- 特性
	- 值类型
	- 值不可变
- 分类
	- 解释字符串
		- 双引号
	- 非解释字符串
		- 反引号
##### <span id="variable">变量</span>
###### 关键词
- var
###### 格式
- var identifier type
###### 特点
- 自动赋值为0值
###### 命名规则
- 骆驼命名法
###### 作用域
- 全局变量
- 局部变量

****
##### <span id="const">常量</span>
###### 关键词
- const
###### 格式
- const identifier [type] = value
- const Pi = 3.14159
###### 要求
- 编译时刻值要确定
- 除了内置函数 例如len()
###### 常数表达式
- 可以执行任意精度的运算
- 没有确定类型
	- 根据上下文需要自动确定类型

****

##### <span id="for">for循环</span>
- for 初始化语句; 条件语句; 修饰语句 {}
- 无限循环  for{}
- for-range
	- for pos, char := range str {}

****
##### <span id="if">if/else判断</span>

- if initialization; condition {
    // do something 
}
- if condition {
    // do something 
} else {
    // do something 
}
- if condition1 {
    // do something 
} else if condition2 {
    // do something else    
} else {
    // catch-all or default
}
- 格式要求
	- 左括号和条件同行
	- 右括号和下一级同行

**** 
##### <span id="witch">switch分支</span>
- switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
- 特性
	- var 可以是任何类型
	- 可以为表达式
	- 各个直接的类型相同
	- 不需要单独使用break
		- fallthrough 关键字接着执行
****

##### <span id="array"> 数组</span>
###### 概念
- 相同唯一类型
	- 原始类型
- 编号
- 长度固定
###### 定义
- var identifier [len]type
### new创建
- var arr=new([5]int)
	- arr类型为 *[5]int

****

##### <span id="slice">切片</span>

- 长度可变的数组
- 计算容量 cap()
- 定义
	- var identifier []type
- append
- copy

****

#### <span id="map">map</span>

###### 概念
- 引用类型
- 初始化
	- 初始化方法
		- map literals
			- maplit = map[string]int{"one":1,"two":2}
		- 引用类型
			- mapcreat:=make(map[string]int)
		- 不要使用new
	- 未初始化
		- nil

- 声明与赋值
	- 要求
		- key
			- 可以：int float  string   指针  接口类型等
			- 可以：结构体
				- key()
				- hash()
			- 不可以：数组  切片 结构体
		- value
			- 任意值
				- 切片（一对多）
					- map1  := make(map[int][]int)
					- map1 := make(map[int]*[]int)
			- 函数
	- 值赋值
		- map1[key] =value
		- val :=map[key]
	- var map map[keytype] valuetype
- 容量
	- 容量不定
	- 动态伸缩
	- 可以在make时候进行指定
		- mapcreated := make(map[string]int,100)
	- 到达容量上限会自动增加
	- 出于性能考虑
		- 尽量表明容量
- 存在判断
	- value,IsPresent = map[key]
- 删除
	- delete(map1,key1)
- 循环
	- for key value := range map{ }
- map 切片
	- 1. 分配切片
		- maplist  := make([] map[string]int,100)
	- 2. 对切片的每个元素进行分配值
		- for i := range mapList{
         mapList[i]=map[string]int{"one":1,"two":2}
	}
- map 排序
	- 无序的
	- 解决方案
		- 复制到切片中
		- 在切片中进行排序 