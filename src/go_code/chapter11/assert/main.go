package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	//b=a 不可以
	//b=a.(Point)
	b = a.(Point)
	fmt.Println(b)
}
