package main

import "fmt"

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-13 16:44:52
 * @LastEditTime: 2021-11-13 23:31:19
 * @LastEditors: Do not edit
 */
func main() {
	var key byte
	fmt.Println("请输入一个字符,a,b,c,d,e,f,g")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a', 'e':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")
	default:
		fmt.Println("输入有误")
	}
}
