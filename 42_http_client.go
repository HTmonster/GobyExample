/*
 * @Author: Theo_hui
 * @Date: 2021-08-25 14:46:05
 * @Descripttion:
 */
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// 发起请求
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//响应信息：
	// 状态
	fmt.Println(resp.Status)
	// body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
}
