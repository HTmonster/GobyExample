/*
 * @Author: Theo_hui
 * @Date: 2021-08-23 11:03:20
 * @Descripttion:
 */
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(sEnc)
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(sDec)

	uEnc := base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(sEnc)
	fmt.Println(uDec)

}
