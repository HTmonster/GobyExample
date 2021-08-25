/*
 * @Author: Theo_hui
 * @Date: 2021-08-25 10:21:52
 * @Descripttion:
 */
package main

import (
	"fmt"
	"testing"
)

//要测试的函数
func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//测试函数
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

//表驱动方式来进行单元测试
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{-1, 2, -1},
		{100, 0, 0},
		{12, -1, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d %d", tt.a, tt.b)

		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Error("got %d, want %d", ans, tt.want)
			}
		})
	}
}
