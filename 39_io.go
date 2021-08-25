/*
 * @Author: Theo_hui
 * @Date: 2021-08-24 09:12:44
 * @Descripttion:
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*==============读取文件==================*/
	//1. ioutil实用方法
	dat, err := ioutil.ReadFile("./README.md")
	check(err)
	fmt.Println(string(dat))

	//2. os方法
	f, err := os.Open("./README.md") //打开文件
	defer f.Close()                  //关闭
	check(err)
	b1 := make([]byte, 3000) //创建读取容器
	n1, err := f.Read(b1)    //读取
	check(err)
	fmt.Println(n1, string(b1))
	o2, err := f.Seek(40, 0) //改变光标位置
	check(err)
	fmt.Println(o2)

	//2-1. bufio包进行缓冲
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Println(string(b4))

	/*===============写入文件=================*/
	//1. ioutil方法
	d1 := []byte("hello\n")
	err = ioutil.WriteFile("./test.txt", d1, 0644)
	check(err)

	//2. os方法
	wf, err := os.Create("./test2.txt") //创建文件
	check(err)
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := wf.Write(d2) //Write写入字节切片
	check(err)
	fmt.Println(n2)
	n3, err := wf.WriteString("hahahahha\n") //WriteString写入字符串
	check(err)
	fmt.Println(n3)
	//2-1 bufio缓冲进行写入
	w := bufio.NewWriter(wf)
	n4, err := w.WriteString("yeyeyeyeyyey\n")
	fmt.Println(n4)
	w.Flush()

	/*======================文件路径=============*/
	p := filepath.Join("dir1", "dir2", "filename.json")
	fmt.Println(p)
	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.Base(p))
	fmt.Println(filepath.IsAbs(p))
	fmt.Println(filepath.Ext(p))

	/*======================目录================*/
	//os.Mkdir()
	//os.MkdirAll()
	//ioutil.ReadDir()
	//os.Chdir()

	/*======================临时文件与目录========*/
	//ioutil.TempFile()
	//ioutil.TempDir()

	/*======================行过滤器==============*/
	scanner := bufio.NewScanner(os.Stdin) //创建扫描器

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
