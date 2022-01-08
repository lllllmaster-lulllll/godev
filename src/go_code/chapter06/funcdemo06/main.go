/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-07 21:47:56
 * @LastEditTime: 2022-01-07 22:00:35
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
)

var (
	//Fun1就是全局匿名函数
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//求两数之和,使用匿名函数完成
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(3, 5)
	fmt.Println("res1", res1)
	//将匿名函数func(n1 int, n2 int) int 赋给a变量
	//则a的数据类型就是函数类型
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(7, 8)
	fmt.Println("res2=", res2)
	res3 := Fun1(4, 5)
	fmt.Println("res3=", res3)
}
