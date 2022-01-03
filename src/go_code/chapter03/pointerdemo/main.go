package main

import "fmt"

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-07 20:50:39
 * @LastEditTime: 2021-11-07 21:00:26
 * @LastEditors: Do not edit
 */
func main() {
	//基本数据类型在内存布局
	var i int = 110
	fmt.Println("i 的地址=", &i)
	var ptr *int = &i
	fmt.Printf("ptr =%v\n", ptr)
	fmt.Println("ptr 的地址=", &ptr)
	fmt.Println("ptr 的地址=", *ptr)
}
