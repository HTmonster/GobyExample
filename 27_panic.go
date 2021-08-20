/*
 * @Author: Theo_hui
 * @Date: 2021-08-20 10:41:40
 * @Descripttion:
 */
package main

import (
	"fmt"
	"os"
)

//创建文件
func createFile(p string) *os.File {
	fmt.Println("[+] creating ", p)

	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

//写文件
func writeFile(s string, f *os.File) {
	fmt.Println("==>writing ", s)
	fmt.Fprintln(f, s)
}

//关闭文件
func closeFile(f *os.File) {
	fmt.Println("[-] closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("test.txt")
	defer closeFile(f)
	writeFile("hahah", f)

}
