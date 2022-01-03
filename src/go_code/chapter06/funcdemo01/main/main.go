/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-16 20:18:12
 * @LastEditTime: 2021-11-17 15:12:30
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
	"go_code/chapter06/funcdemo01/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.5
	var operator byte = '/'
	res := utils.Cal(n1, n2, operator)
	fmt.Println("res=", res)
}
