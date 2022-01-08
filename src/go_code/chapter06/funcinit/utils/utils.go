/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-06 22:53:02
 * @LastEditTime: 2022-01-06 22:55:29
 * @LastEditors: Do not edit
 */
package utils

var Age int
var Name string

//Age和Name全局变量,需要在main.go中使用
//但是需要初始化

//init 函数完成初始化工作
func init() {
	Age = 100
	Name = "tom"
}
