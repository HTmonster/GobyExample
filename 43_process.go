/*
 * @Author: Theo_hui
 * @Date: 2021-08-25 15:21:27
 * @Descripttion:
 */
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	/*====================创建进程==============*/
	//创建进程
	ipconfigCmd := exec.Command("ipconfig")
	//执行并获得输出
	out, err := ipconfigCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	//因为操作系统不一样，剩余部分略过
}
