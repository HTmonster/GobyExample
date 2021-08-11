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
- 高级数据结构
	- [x] array
	- [x] slice
	- [x] map
	- [x] :adhesive_bandage: range
- 函数
	- [x] 普通
	- [x] 多值返回
	- [x] 变参函数
	- [x] 迭代
- 派生类型
	- [x] 指针
	- [x] 结构体
	- [x] 方法
	- [x] 接口
- 错误处理
	- [x] 错误处理
- 协程与通道
	- [x] 协程
	- [x] 通道
		- [x] 通道缓冲
		- [x] 通道同步
		- [x] 通道方向
		- [x] 通道选择器
		- [x] 超时处理
		- [x] 非阻塞通道
		- [x] 通道关闭
		- [x] 通道的遍历
	- [x] Timer&Ticker



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
| Golang range|1.可以只迭代map的key 2.在字符串中迭代 unicode||
| Golang 函数||[函数](#function)|
| Golang 指针|参考C|[指针](#pointer)|
| Golang struct||[struct](#struct)|
| Golang 方法|Go会自动处理值与指针之间的转换|[方法](#method)|
| Golang 接口|1. 实现多类 2.intetface{}||
| Golang 错误处理||[错误处理](#error)|
| Golang 协程||[协程](#goroutine)|
| Golang 通道||[通道](#channels)|
| Golang 定时器与打点器||[Timer&Ticker](#Timer)|



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
###### new创建
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

****
#### <span id="function"> 函数 </span>

- 格式 
	- func functionName(parameter_list) (return_value_list) {
   …
}


- main函数
	- 无参数
	- 无返回值 

- 可以多值返回
- 变参函数
	- ...
- 闭包，匿名函数
- 递归

****
#### <span id="pointer"> 指针</span>
###### 控制指定集合
- 数据结构
- 分配的数量
- 内存访问模式
###### 取地址
- &

###### 空指针
- nil

###### 不能发得到一个文字 或者常量的地址

****
#### <span id="struct">结构体</span>
###### 结构体
- 字段构成
	- 名字（唯一）
	- 类型
###### 定义

- type identifier struct {
    field1 type1
    field2 type2
    ...
}
- 支持匿名字段
	- type identifier struct {
    field1 type1
    type2
    ...
	}
- 支持内嵌结构体
	- type A struct {
    	a  int
	} 
	type B struct {
    	s  str
    	A
	}

###### 初始化
- 混合字面量语法

	- ms := &struct1{10, 15.5, "Chris"}
	- intr := Interval{0, 3}            (A)
intr := Interval{end:5, start:1}  (B)
intr := Interval{end:5}           (C)

###### 字段赋值
- var s T
s.a = 5
s.b = 8
- 无论是值或者指针，都可以使用选择器（.)来进行访问

###### new函数

- 为结构体分配内存并返回指针
	- var t *T = new(T)

###### 使用工厂方法创建实例

###### 标签（可选）
- 附属于字段的字符串
- 使用reflect包进行访问
	- 在运行时自省类型、属性和方法

****
#### <span id="method"> 方法</span>
###### 简介
- 作用在接收者（receiver) 上
- 特殊函数
###### 定义
- func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
- 接收者
	- 除了接口类型之外都可以
- 不使用recv 的值，可以用 _ 替换它
###### 方法与函数的区别
- 变量
	- 方法被变量调用
	- 函数将变量作为参数

###### 可以用来定制String() 方法

****
#### <span id="error"> 错误处理</span>
- 使用一个明确的返回值来传递错误信息
	- 通常是最后一个返回值
	- error类型
	- 内建接口
- `errors.New()`构造一个基本的error值
- `nil`代表没有错误
- 实现`Error()`方法来自定义error类型

****
#### <span id="goroutine">协程</span>

- 启动
	- 在调用<u>函数</u>或者<u>方法</u>时，使用关键字**go**
- 协程的调用会**立即返回**
- `main goroutine` 会运行任何其他 `goroutine`，如果 `main goroutine` 终止，则该程序将被终止，并且其他 `goroutine` 将得不到运行机会
- 常常与**匿名函数**搭配
- M:N线程模型，M个协程运行在N个线程之上

****
#### <span id="channels">通道</span>
- 创建
	- make(chan val-type)
- 发送
	- channel <-
- 接收
	- <-channel 
- 发送和接收默认是**阻塞式**的
- 通道 默认是 **无缓冲**
	- 只有通道准备好接收时，才进行发送
- **有缓冲**的通过
	- make(chan string,2) 指定缓冲两个值
- 可以使用通道来**同步**协程之间的执行状态
- 使用通道作为函数的参数时，可以指定通道是**只读**还是**只写**
- **选择器**可以同时等待多个通道操作
- 超时处理
	- `Time.After()`
- **非阻塞**通道
	- 带 `default`子句的`select`语句实现
- 通道**关闭**
	- `close()` 语句
- **遍历**通道
	- `for range` 

****

#### <span id="timer">Timer定时器与Ticker打点器</span>

- Timer 定时器
	- `func After(d Duration) <- chan Time`
		- 经过时长d之后，向通道Time发送当前消息
		- 返回只读通道
	- `func AfterFunc(d Duration, f func()) *Timer`
		- 在经过时长d之后调用函数f
		- 返回Timer
	- `func NewTimer(d Duration) *Timer`
		- 创建新的定时器
		- `func (t *Timer) Reset(d Duration) bool`
			- 重置定时器，以d为触发时间
		- `func (t *Timer) Stop() bool`
			- 停止定时器
		- `Timer.C` 定时向信道发送时间
- Ticker 打点器
	- `func Tick(d Duration) <-chan Time`
		- 每经过d之后，向通道发送当前时间
	- `func NewTick(d Duration) *Ticker`
		- 创建新的打点器
		- `func (t *Ticker) Stop()`
			- 停止打点器
		- `Ticker.C` 定时向通道发送当前时间