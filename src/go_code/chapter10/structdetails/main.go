package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"` //就是struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	//创建一个monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}
	//2.将monster变量序列化为 字符json格式的字符串
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
