/*
 * @Author: Theo_hui
 * @Date: 2021-08-25 10:51:09
 * @Descripttion:
 */
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	/*================参数============*/
	fmt.Println(os.Args)
	/*===============标志=============*/
	// 声明
	strp := flag.String("name", "rose", "person name")
	nump := flag.Int("age", 0, "person age")
	stup := flag.Bool("isStudent", true, "Is person a student")

	//解析
	flag.Parse()
	fmt.Println(*strp, *nump, *stup)
	/*===============环境变量==========*/
	fmt.Println(os.Environ())

}
