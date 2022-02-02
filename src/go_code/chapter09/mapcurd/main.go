package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	cities["no3"] = "上海~"
	fmt.Println(cities)
	//演示删除
	delete(cities, "no1")
	fmt.Println(cities)
	//当delete指定的key不存在时,删除不会操作,也不会报错
	delete(cities, "no4")
	fmt.Println(cities)
	//演示查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no2  key 值为%v\n", val)
	} else {
		fmt.Println("没有no1")
	}
	// 1）如果我们要删除map的所有key，没有一个专门的方法一次删除，可以遍历一下key，逐个删除
	// 2）或者 map=make（...），make一个新的，让原来的成为垃圾，被gc回收
	cities = make(map[string]string)
	fmt.Println(cities)
}
