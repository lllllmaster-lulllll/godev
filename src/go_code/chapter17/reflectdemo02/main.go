package main

import (
	"fmt"
	"reflect"
)

//通过反射,修改 num int 的值  修改 student 的值

func reflect01(b interface{}) {
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(20)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	// fmt.Println(rVal)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println(num)
}
