/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-18 22:08:38
 * @LastEditTime: 2021-11-18 22:25:38
 * @LastEditors: Do not edit
 */
package main

import "fmt"

func monkeypeach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入天数不对")
	}
	if n == 10 {
		return 1
	} else {
		return (monkeypeach(n+1) + 1) * 2
	}
}
func main() {
	fmt.Println(monkeypeach(9))

}
