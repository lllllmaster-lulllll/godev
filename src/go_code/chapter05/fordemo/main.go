package main

import "fmt"

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-14 15:56:44
 * @LastEditTime: 2021-11-14 16:22:44
 * @LastEditors: Do not edit
 */
func main() {
	// var str string = "hello,world!北京"
	// // str2 := []rune(str)
	// // for i := 0; i < len(str2); i++ {
	// // 	fmt.Printf("%c\n", str2[i])
	// // }

	// for index, val := range str {
	// 	fmt.Printf("index=%d,val=%c \n", index, val)
	// }
	var count int = 0
	var sum int = 0
	for i := 1; i <= 100; i++ {

		if i%9 == 0 {
			count++
			sum += i
			fmt.Println(i)
		}
	}
	fmt.Println("1-100之间所有 9 的倍数的个数为: ", count)
	fmt.Println("1-100之间所有 9 的倍数的总和为: ", sum)
}
