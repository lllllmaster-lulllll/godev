/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-16 20:51:12
 * @LastEditTime: 2021-11-16 20:51:12
 * @LastEditors: Do not edit
 */
package utils

import "fmt"

func Cal(n1 float64, n2 float64, operator byte) float64 {

	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Printf("操作符号错误...")
	}
	return res
}
