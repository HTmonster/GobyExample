/*
 * @Author: Theo_hui
 * @Date: 2021-08-23 10:47:31
 * @Descripttion:
 */
package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	if u, err := url.Parse(s); err == nil {
		fmt.Println(u.Scheme)
		fmt.Println(u.User)
		fmt.Println(u.User.Username())
		fmt.Println(u.User.Password())
		fmt.Println(u.Host)
		fmt.Println(u.Path)
		fmt.Println(u.Fragment)
		fmt.Println(u.RawQuery)
	}
}
