package main

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("第 %v 个参数是 bool 类型,值是 %v\n", index+1, v)
		case float32:
			fmt.Printf("第 %v 个参数是 float32 类型,值是 %v\n", index+1, v)
		case float64:
			fmt.Printf("第 %v 个参数是 float64 类型,值是 %v\n", index+1, v)
		case int, int32, int64:
			fmt.Printf("第 %v 个参数是 整数 类型,值是 %v\n", index+1, v)
		case string:
			fmt.Printf("第 %v 个参数是 string 类型,值是 %v\n", index+1, v)
		case Student:
			fmt.Printf("第 %v 个参数是 Student 类型,值是 %v\n", index+1, v)
		case *Student:
			fmt.Printf("第 %v 个参数是 *Student 类型,值是 %v\n", index+1, v)
		default:
			fmt.Printf("第 %v 个参数类型不确定,值是 %v\n", index+1, v)
		}
	}
}

func main() {
	var n1 float32 = 1.2
	var n2 float64 = 2.3
	var n3 int32 = 789
	var name string = "Tom"
	address := "北京"
	n4 := 300
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)

}
