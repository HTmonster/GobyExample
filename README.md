#### :rocket: GobyExample

> 跟随 [Go by Example 中文版 (gobyexample-cn.github.io)](https://gobyexample-cn.github.io/) 项目复习Golang语言



#### :bicyclist: 完成清单

- [x] Hello world
- [x] value
- [x] variable
- [x] constant



#### :warning: 注意点

| 知识点 | 注意内容 | 链接 |
| ------ | -------- | ---- |
|   Golang值基本类型     |          |  [基本类型](#value)|
|   Golang变量     |          |  [变量](#variable)   |
|      Golang常量  |          | [常量](#const)     |

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
