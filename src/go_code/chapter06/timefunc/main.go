/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2022-01-10 00:06:52
 * @LastEditTime: 2022-01-10 13:48:18
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	//看看日期和时间相关的函数
	//1.获取当前时间
	now := time.Now()

	fmt.Printf("now=%v now type=%T\n", now, now)
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month())) //fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())
	//格式化时间
	//方式一:
	fmt.Printf("当前年月日 %d-%d-%d  %d-%d-%d  \n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println("-----------------------")
	//方式二
	//now.Format("2006/01/02 15:04:05")也可以now.Format("2006-01-02 15:04:05")
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Println(now.Format("15:04:05"))
	fmt.Println()

	//时间的常量
	//需求,每隔一秒钟打印一个数字,大隐刀100时退出

	// i := 0

	// for {
	// 	i++
	// 	fmt.Println(i)
	// 	// time.Sleep(time.Second)
	// }
	// //需求2,每隔一秒钟打印一个数字,大隐刀100时退出
	// j := 0
	// for {
	// 	j++
	// 	fmt.Println(time.Millisecond * 100)
	// }

	//Unix和UnixNano的使用
	fmt.Printf("unix时间戳=%v  UnixNano时间戳=%v", now.Unix(), now.UnixNano())

}
