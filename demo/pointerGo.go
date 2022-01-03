package main

import "fmt"

const MAX int = 3
func main() {
	//var a int = 10
	//fmt.Printf("变量的地址：%x\n",&a)
	//var a int=20
	//var ip *int
	//ip=&a
	//fmt.Printf("a 变量的地址是：%x\n",&a)
	//fmt.Printf("ip 变量的地址是：%x\n",ip)
	//fmt.Printf("*ip 变量的地址是：%d\n",*ip)
	//var ptr *int
	//fmt.Printf("ptr 的值为：%x\n",ptr)
	//
	a:=[]int{10,100,200}
	var i int
	var ptr [MAX]*int;

	for i=0;i<MAX;i++{
		ptr[i] = &a[i]
	}
	for i=0;i<MAX;i++{
		fmt.Printf("a[%d] = %d\n", i,*ptr[i])
	}
}
