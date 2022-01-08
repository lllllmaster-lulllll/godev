/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-07 22:20:58
 * @LastEditTime: 2022-01-08 15:37:43
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
	"strings"
)

//累加器函数
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

// 1)编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg)，则返回 文件名.jpg ，如果
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix ，该函数可以判断某个字符串是否有指定的后缀。

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果name没有指定后缀,则加上,否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	//累加器
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	//测试makesuffix的使用
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后为:", f2("winter"))
	fmt.Println("文件名处理后为:", f2("kangnan.jpg"))
}
