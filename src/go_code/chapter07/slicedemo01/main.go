package main

import "fmt"

func main() {
	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//声明/定义一个切片
	//slice:=intArr[1:3]----从1开始到3(但是不包含3)
	//1. slice就是切片名
	//2. intArr[1:3]表示slice引用到intArr这个数组的第二个元素到第三个元素
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 元素是 的=", slice)
	fmt.Println("slice 的元素长度=", len(slice))
	fmt.Println("slice 的容量=", cap(slice)) //切片的容量是可以动态变化的capability
	fmt.Printf("intArr[1]的地址%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址%p\n", &slice[0])
	slice[1] = 67890
	fmt.Println("intArr=", intArr)

}
