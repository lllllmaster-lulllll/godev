/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-03 20:08:31
 * @LastEditTime: 2022-01-03 22:05:27
 * @LastEditors: Do not edit
 */
package main

import "fmt"

type myFunType func(int, int) int

//在Go中,函数也是一种数据类型
//可以复制给一个变量,则该变量就是一个函数类型的变量了.通过该变量对函数调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}
func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return sum, sub
}
func main() {
	a := getSum
	fmt.Printf("a的类型%T,getSum类型是%T\n", a, getSum)
	res := a(10, 40)
	fmt.Println("res=", res)
	//看案例
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)
	//给int取了别名,在go中myInt和int 虽然都是int类型,但是go认为myInt和int是两个类型
	type myInt int
	var num1 myInt
	num1 = 40
	var num2 int
	num2 = int(num1)
	fmt.Println("num1=", num1, "num2=", num2)

	//再加一个案例

	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res3=", res3)

	//看案例
	//支持返回值重命名
	sumRes, subRes := getSumAndSub(1, 2)
	fmt.Printf("a=%v,b=%v\n", sumRes, subRes)
}
