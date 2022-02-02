package main

import (
	"fmt"
)

func main() {
	//map的声明和注意事项
	var a map[string]string
	//在使用map前,需要先make,make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "曹操"
	a["no2"] = "关羽"
	a["no1"] = "诸葛亮"
	fmt.Println(a)

}
