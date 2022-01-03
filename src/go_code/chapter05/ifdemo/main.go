package main

import "fmt"

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-13 16:44:52
 * @LastEditTime: 2021-11-13 17:11:26
 * @LastEditors: Do not edit
 */
func main() {
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("成年")
	}

}
