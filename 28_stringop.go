/*
 * @Author: Theo_hui
 * @Date: 2021-08-20 11:13:23
 * @Descripttion:
 */
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Contains:	", strings.Contains("pineapple", "apple"))
	fmt.Println("Count:		", strings.Count("pineapple", "p"))
	fmt.Println("HasPrefix:	", strings.HasPrefix("pineapple", "pi"))
	fmt.Println("HasSuffix:	", strings.HasSuffix("pineapple", "apple"))
	fmt.Println("Index:		", strings.Index("pineapple", "e"))
	fmt.Println("Join:		", strings.Join([]string{"p", "a"}, "a"))
	fmt.Println("Repeat:	", strings.Repeat("pi", 3))
	fmt.Println("Replace:	", strings.Replace("pineapple", "p", "q", -1))
	fmt.Println("Replace:	", strings.Replace("pineapple", "p", "q", 1))
	fmt.Println("Split:		", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:	", strings.ToLower("Apple"))
	fmt.Println("ToUpper:	", strings.ToUpper("bubu"))
}
