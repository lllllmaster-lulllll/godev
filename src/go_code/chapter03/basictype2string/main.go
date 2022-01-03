package main

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-07 16:37:31
 * @LastEditTime: 2021-11-07 17:10:49
 * @LastEditors: Do not edit
 */

import (
	"fmt"
	"strconv"
	_ "unsafe"
)

func main() {
	var num1 int = 99
	var num2 float64 = 13.456
	var b bool = true
	var myChar byte = 'h'
	var str string //空的字符串
	//使用第一种方式来转换 fmt.Sprintf 方法

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)

	//第二种方式转换
	var num3 int = 99
	var num4 float64 = 45.678
	var b2 bool = true
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str = %q \n", str, str)
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str = %q \n", str, str)
	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str = %q \n", str, str)
	//strconv 包中有一个函数 itoa
	var num5 int = 'a'
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str = %v \n", str, str)
}
