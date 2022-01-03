/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-18 21:45:36
 * @LastEditTime: 2021-11-18 22:02:26
 * @LastEditors: Do not edit
 */
package main

import "fmt"

func fb(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fb(n-1) + fb(n-2)
	}
}

func main() {
	var n int
	fmt.Print("请输入想要想要的数字: _")
	fmt.Scanln(&n)
	sum := fb(n)
	fmt.Println("sum=", sum)

}
