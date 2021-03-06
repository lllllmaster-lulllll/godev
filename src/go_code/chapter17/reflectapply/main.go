package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"monster_name"`
	Age   int     `json:"monster_age"`
	Score float64 `json:"monster_score"`
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("----------start----------")
	fmt.Println(s)
	fmt.Println("-----------end-----------")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// func (s Monster) Set(name string, age int, score float64, sex string) {
// 	s.Name = name
// 	s.Age = age
// 	s.Score = score
// 	s.Sex = sex
// }
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
		numofMethod := val.NumMethod()
		fmt.Printf("struct has %d methods\n", numofMethod)
		//var params []reflect.Value
		//函数排序按照名字字母排序,A->Z
		val.Method(1).Call(nil)
		//调用结构体的第1个方法Method(0)
		var params []reflect.Value
		params = append(params, reflect.ValueOf(10))
		params = append(params, reflect.ValueOf(40))
		res := val.Method(0).Call(params) //传入的参数是[]reflect.Value
		fmt.Println("res=", res[0].Int()) //返回结果,返回的结果是 []reflect.Value
	}

}

func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(a)
}
