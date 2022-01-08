/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-06 22:42:59
 * @LastEditTime: 2022-01-07 18:32:36
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
	"go_code/chapter06/funcinit/utils"
)

//为了看到全局变量是被先执行的,先写函数
var age = test()

func test() int {
	fmt.Println("test()")
	return 90
}

//init  通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init.....")
}

//main函数
func main() {
	age = age
	fmt.Println("main.....")
	fmt.Println("Age=", utils.Age, " Name=", utils.Name)
}
