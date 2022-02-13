package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{}) {
	//通过反射获取传入变量的ype kind value
	//先获取 reflect.Type()
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//进行加减
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Println("rVal=", rVal)
	//将 rVal 转成 interface{}
	iV := rVal.Interface()
	//将 interface 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)

}

//专门演示对结构体的反射
func reflectTest02(b interface{}) {
	//通过反射获取传入变量的ype kind value
	//先获取 reflect.Type()
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//获取变量对应的 Kind
	//1. rVal.Kind()==>
	kind1 := rVal.Kind()
	//2. rType.Kind()==>
	kind2 := rType.Kind()

	fmt.Printf("kind1=%v kind2=%v\n", kind1, kind2)
	//将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV = %v iV = %T\n", iV, iV)
	//将 interface{}通过断言转成需要的类型
	//简单使用类型检查
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//请编写一个案例演示对(基本数据类型、 interface{}、reflect.Value)进行反射的基本操作

	// var num int = 100
	// reflectTest01(num)

	//定义一个 Student 的实例
	stu := Student{
		Name: "Tom",
		Age:  8,
	}
	reflectTest02(stu)
}
