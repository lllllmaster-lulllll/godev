package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

func main() {
	//创建要给Student实例
	// var stu = model.student{
	// 	Name:  "kangnan",
	// 	Score: 78,
	// }
	//当student结构体是首字母小写,我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom~", 88.8)
	fmt.Println(*stu) //&{....}

	fmt.Println(stu.GetScore())

}
