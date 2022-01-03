package main

import "fmt"

/*
 * @Description:
 * @version:1.0
 * @Author: Nan Kang
 * @Date: 2021-11-11 00:02:39
 * @LastEditTime: 2021-11-13 16:30:07
 * @LastEditors: Do not edit
 */

func main() {
	var i int = 5
	fmt.Printf("i=%b\n", i)
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)
	fmt.Println(-2 ^ 2)
}
