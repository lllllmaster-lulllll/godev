package main

import (
	"fmt"
)

//指针，slice，和map的零值都是 nil，即还没有分配空间。
//如果需要使用这样的字段,需要先make才能使用
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}

func main() {
	//定义结构体变量
	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("ok ptr")
	}
	if p1.slice == nil {
		fmt.Println("ok slice")
	}
	if p1.map1 == nil {
		fmt.Println("ok map1")
	}

}
