package main

import (
	"fmt"
)

//定义一个cat结构体,将CAT的各个字段/属性信息,放入到cat结构体进行管理
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

func main() {
	//使用struct来完成案例
	//创建一个Cat的变量
	var cat1 Cat //var a int
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃 <·)))><<"
	fmt.Println("cat1=", cat1)
	fmt.Println("猫猫的信息如下: ")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("hobby=", cat1.Hobby)
}
