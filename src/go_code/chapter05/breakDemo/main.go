/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-14 20:58:38
 * @LastEditTime: 2021-11-16 20:17:16
 * @LastEditors: Do not edit
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 0
	for {
		//为随机数设置一个种子
		rand.Seed(time.Now().UnixNano())
		//随机生成 1-100 整数
		n := rand.Intn(101) //[0,101)
		fmt.Println(n)
		count++
		if n == 99 {
			break
		}
		time.Sleep(1000)
	}
	fmt.Println("生成 99 一共使用了 ", count, " 次")

}
