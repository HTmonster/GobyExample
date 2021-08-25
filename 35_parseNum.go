/*
 * @Author: Theo_hui
 * @Date: 2021-08-23 10:37:38
 * @Descripttion:
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("1234", 0, 64)
	fmt.Println(i)

	x, _ := strconv.ParseInt("0x1234", 0, 64)
	fmt.Println(x)

	k, _ := strconv.Atoi("23")
	fmt.Println(k)
}
