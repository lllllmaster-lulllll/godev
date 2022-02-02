package main

import (
	"fmt"
)

func main() {
	// 要求：使用一个map来记录monster的信息name和age，
	// 也就是说一个monster对应一个map并且妖怪的个数可以动态的增加=>map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"
	}
	//数组越界
	// if monsters[2] == nil {
	// 	monsters[2] = make(map[string]string, 2)
	// 	monsters[2]["name"] = "ksd"
	// 	monsters[2]["age"] = "400"
	// }
	//使用到切片的append函数,可以动态增加数据
	newMonster := map[string]string{
		"name": "新的妖怪",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
