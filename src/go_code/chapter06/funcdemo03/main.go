/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-17 16:49:50
 * @LastEditTime: 2021-11-17 16:52:07
 * @LastEditors: Do not edit
 */
package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n= ", n)
}

func main() {
	test(4)
}
