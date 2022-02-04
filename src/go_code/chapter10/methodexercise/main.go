package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}
type Integer int

func (i Integer) print() {
	fmt.Println("i=", i)
}

func (circle Circle) area() float64 {
	return circle.radius * circle.radius * math.Pi
}

type Student struct {
	Name string
	Age  int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v],Age=[%v]", stu.Name, stu.Age)
	return str
}
func main() {
	//定义一个Student变量
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(stu)  //按默认方式输出
	fmt.Println(&stu) //按新实现的方式输出
}
