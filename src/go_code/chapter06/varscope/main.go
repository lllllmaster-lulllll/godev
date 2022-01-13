/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-08 22:06:01
 * @LastEditTime: 2022-01-08 22:57:28
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
)

var name = "tom"

func test01() {
	fmt.Println(name)
}
func test02() {
	name = "jack"
	fmt.Println(name)
}
func main() {
	// fmt.Println("age=", age)
	// fmt.Println("Name=", Name)
	fmt.Println(name)
	test01()
	test02()
	test01()

}
