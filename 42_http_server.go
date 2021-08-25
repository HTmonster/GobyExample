/*
 * @Author: Theo_hui
 * @Date: 2021-08-25 14:52:58
 * @Descripttion:
 */
package main

import (
	"fmt"
	"net/http"
	"time"
)

//1.简单的handler
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

//2.略微复杂的handler
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v:%v\n", name, h)
		}
	}
}

//3.带context的handler
func context(w http.ResponseWriter, req *http.Request){
	//创建一个context
	ctx := req.Context()
	fmt.Println("context: handler started")
	def fmt.Println("context: handler ended")

	select{
		//等待一段时间，模拟服务器在哦你住
	case <-time.After(10*time.Second):
		fmt.Fprintf(w,"context\n")
		//通道发送信号，取消工作尽快返回
	case <-ctx.Done():
		err := ctx.Err() //返回通道关闭的原因
		fmt.Println(err)
		internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// 设置路由
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/context",context)

	//监听
	http.ListenAndServe(":8080", nil)

}
