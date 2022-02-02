package main

import "unsafe"

const(
	a="abc"
	b=len(a)
	c=unsafe.Sizeof(a)
)
func main() {
	//const LENGHT int = 10
	//const WIDTH int=5
	//var area int
	//const a,b,c =1, false, "str" //多重赋值
	//area=LENGHT*WIDTH
	//fmt.Printf("面积为：%d", area)
	//println()
	//println(a,b,c)
	println(a,b,c)
}

