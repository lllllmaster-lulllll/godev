package main

import "fmt"

/*
 * @Description:
 * @version:
 * @Author: Nan Kang
 * @Date: 2021-11-10 23:36:42
 * @LastEditTime: 2021-11-10 23:56:34
 * @LastEditors: Do not edit
 */

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	// fmt.Println("请输入姓名:")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪水:")
	// fmt.Scanln(&sal)
	// fmt.Println("请输入是否通过考试:")
	// fmt.Scanln(&isPass)

	// fmt.Printf("姓名: %v\n 年龄:%v\n 薪水:%v\n 考试情况:%v", name, age, sal, isPass)
	fmt.Println("请输入姓名,年龄,薪水和考试情况")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名: %v\n 年龄:%v\n 薪水:%v\n 考试情况:%v\n", name, age, sal, isPass)
}
