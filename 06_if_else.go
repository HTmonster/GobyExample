/*
 * @Author: Theo_hui
 * @Date: 2021-07-28 14:42:05
 * @Descripttion:
 */
package main

import "fmt"

func main() {

	if true {
		fmt.Println("it is true")
	}

	if false {
		fmt.Println("it is true")
	} else {
		fmt.Println("it is false")
	}

	i := 3
	if i == 1 {
		fmt.Println("it is 1")
	} else if i == 2 {
		fmt.Println("it is 2")
	} else if i == 3 {
		fmt.Println("it is 3")
	} else {
		fmt.Println("it is bigger than 3")
	}
}
