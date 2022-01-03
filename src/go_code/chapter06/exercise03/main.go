/*
 * @Description:
 * @version:1.0
 * @Author: Nan Kang
 * @Date: 2022-01-03 17:00:18
 * @LastEditTime: 2022-01-03 17:12:56
 * @LastEditors: Do not edit
 */
package main

import "fmt"

//n的范围是 1--10之间
func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对!!!")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
func main() {
	fmt.Println("第9天的桃子数量是=>", peach(9))
}
