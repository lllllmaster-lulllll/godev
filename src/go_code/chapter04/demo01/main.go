package main

import "fmt"

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-07 22:15:10
 * @LastEditTime: 2021-11-10 22:10:51

 */
func main() {
	//在 golang 中,++和--只能独立使用
	var i int = 8
	var a int = 9
	a, i = i, a
	fmt.Println("a=", a, " i=", i)
	i++
	a = i
	fmt.Println(a)

}
