/*
 * @Author: Theo_hui
 * @Date: 2021-08-23 10:26:15
 * @Descripttion:
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())

	rand.New(rand.NewSource(time.Now().UnixNano()))
}
