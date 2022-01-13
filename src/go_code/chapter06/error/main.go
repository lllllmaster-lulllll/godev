package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		err := recover() //recover()内置函数,可以捕获到异常
		//if err := recover();err != nil{
		// code
		// }//与下面的写法等价
		if err != nil { //说明捕获到异常
			fmt.Println("err=", err)
			//发送错误信息给管理员
			fmt.Println("发送邮件给admin@sohu.com") //伪代码
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//函数去读取以配置文件init.conf的信息
//如果文件名传入不正确,我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}
func test02() {
	err := readConf("ssconfig.ini")
	if err != nil {
		//如果读取文件发送错误,就输出这个错误,并终止程序
		panic(err)
	}
	fmt.Println("test02()继续执行")
}
func main() {
	// test()
	// fmt.Println("main()下面的代码...")
	//测试自定义错误的使用
	test02()
	fmt.Println("main()继续执行")
}
