package main

import "fmt"

func main() {
	str := "hello@atguigu康安"
	slice := str[6:]
	fmt.Println("slice=", slice)
	//str[0] = 'z'编译不通过
	//"hello@atguigu"=>改成"zello@atguigu""
	arr1 := []byte(str)
	for i, v := range arr1 {
		fmt.Printf("第%v个字符为%c\n", i, v)
	}
	arr2 := []rune(str)
	arr2[0] = '即'
	for i, v := range arr2 {
		fmt.Printf("第%v个字符为%c\n", i, v)
	}
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

}
