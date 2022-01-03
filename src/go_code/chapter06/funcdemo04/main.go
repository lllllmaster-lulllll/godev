/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-03 19:30:27
 * @LastEditTime: 2022-01-03 19:44:27
 * @LastEditors: Do not edit
 */
package main

import "fmt"

func test() {
	//n1是test函数的局部变量,只能在test函数中使用
	// var n1 int = 10

}
func test02(n1 int) {
	n1 = n1 + 10
	fmt.Println("test02 n1=", n1)
}
func test03(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test03 n1=", *n1)
}
func main() {
	// n1 := 20
	// test02(n1)
	// fmt.Println("main n1=", n1)
	num := 20
	test03(&num)
	fmt.Println("main()  num= ", num)
}
