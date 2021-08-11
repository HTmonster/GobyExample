/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 16:28:58
 * @Descripttion:
 */
package main

import "fmt"

func main() {
	/* 切片创建 */
	// 1. 一般式 不初始化
	var slice_a []int
	fmt.Printf("value: %v \texpress:%#v \tlen:%d \tcap:%d\n", slice_a, slice_a, len(slice_a), cap(slice_a))

	// 2. 一般式 初始化
	slice_b := []int{2, 0, 2, 0}
	fmt.Printf("value: %v \texpress:%#v \tlen:%d \tcap:%d\n", slice_b, slice_b, len(slice_b), cap(slice_b))

	// 3. 通过make
	slice_c := make([]float32, 2, 4)
	fmt.Printf("value: %v \texpress:%#v \tlen:%d \tcap:%d\n", slice_c, slice_c, len(slice_c), cap(slice_c))
	slice_c[0] = 2.15
	fmt.Printf("value: %v \texpress:%#v \tlen:%d \tcap:%d\n", slice_c, slice_c, len(slice_c), cap(slice_c))

	/* 操作 */
	// copy操作
	slice_big := make([]int, 8)
	copy(slice_big, slice_b)
	fmt.Printf("value: %v \texpress:%#v \tlen:%d \tcap:%d\n", slice_big, slice_big, len(slice_big), cap(slice_big))

	// append操作
	slice_big = append(slice_big, 1, 2, 3)
	fmt.Printf("value: %v \texpress:%#v \tlen:%d \tcap:%d\n", slice_big, slice_big, len(slice_big), cap(slice_big))

	slice_a = append(slice_a, slice_b...)
	fmt.Printf("value: %v \texpress:%#v \tlen:%d \tcap:%d\n", slice_a, slice_a, len(slice_a), cap(slice_a))

}
