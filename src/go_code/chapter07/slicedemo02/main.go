package main

import (
	"fmt"
)

func main() {

	//方式2
	//演示切片的使用 make
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 89

	//对于切片,必须使用make
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//方式3
	var slice2 []string = []string{"tom", "mary", "jack"}
	fmt.Println("slice2=", slice2)
	fmt.Println("len(slice2)=", len(slice2))
	fmt.Println("cap(slice2)=", cap(slice2))
}
