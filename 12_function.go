/*
 * @Author: Theo_hui
 * @Date: 2021-07-31 10:02:27
 * @Descripttion:
 */
package main

import "fmt"

//func 函数名（参数）（返回值）
func sum(a int, b int) int {

	return a + b
}

//多值返回
func getInfo(key string) (string, int) {
	kvs := map[string]int{"joh": 123, "hank": 234}

	v := kvs[key]

	return key, v
}

// 变参函数
func h_sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

// 闭包，匿名函数
func intSeq() func() int {
	i := 0

	//函数返回一个匿名函数
	//隐藏变量i
	return func() int {
		i++
		return i
	}
}

// 递归函数
func fibo(a int) int {

	if a <= 0 {
		return 1
	} else if a == 1 {
		return 1
	} else {
		return fibo(a-1) + fibo(a-2)
	}
}

func main() {
	// 普通函数
	fmt.Println(sum(1, 2))
	// 多值返回
	fmt.Println(getInfo("joh"))
	// 变参函数
	fmt.Println(h_sum(1, 2, 3, 4, 5, 6))
	nums := []int{33, 44, 55}
	fmt.Println(h_sum(nums...))
	//闭包与匿名函数
	nextInt := intSeq()
	fmt.Println(nextInt)
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//迭代
	fmt.Println(fibo(1))
	fmt.Println(fibo(7))
}
