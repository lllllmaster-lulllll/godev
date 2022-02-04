package main

import (
	"fmt"
)

type MethodUtils struct {
	//字段
}

//给MethodUtils编写方法
func (m MethodUtils) Print(x int, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main() {
	//1）编写结构体(MethodUtils)，编程一个方法，方法不需要参数，
	//在方法中打印一个 10*8 的矩形， 在 main 方法中调用该方法。
	var m MethodUtils
	(&m).Print(5, 6)

}
