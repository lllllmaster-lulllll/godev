package main

import (
	"fmt"
)

func main() {
	//方式1
	//声明,这时map为nil
	var cities map[string]string
	//make(map[string]string,10) 分配一个空map
	cities = make(map[string]string, 10)
	cities["key11"] = "4567"
	fmt.Println(cities)
	//方式2
	//声明,就直接make
	var cities2 = make(map[string]string)
	cities2["key22"] = "6789"
	fmt.Println(cities2)
	//方式3
	//声明,直接赋值
	var cities3 map[string]string = map[string]string{"no4": "成都"}
	cities3["no1"] = "北京"
	fmt.Println(cities3)

	//存储三个学生信息 学生信息包括name和sex
	var studentInfo = make(map[string]map[string]string, 10)
	studentInfo["st01"] = make(map[string]string, 2)
	studentInfo["st01"]["name"] = "tom"
	studentInfo["st01"]["sex"] = "boy"

	studentInfo["st02"] = make(map[string]string, 2)
	studentInfo["st02"]["name"] = "mary"
	studentInfo["st02"]["sex"] = "girl"

	studentInfo["st03"] = make(map[string]string, 2)
	studentInfo["st03"]["name"] = "ss"
	studentInfo["st03"]["sex"] = "boy"
	fmt.Println(studentInfo)
}
