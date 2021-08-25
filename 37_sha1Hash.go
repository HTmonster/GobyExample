/*
 * @Author: Theo_hui
 * @Date: 2021-08-23 10:57:01
 * @Descripttion:
 */
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	fmt.Println(h)

	h.Write([]byte("a string"))
	hs := h.Sum(nil)
	fmt.Println(hs)
}
