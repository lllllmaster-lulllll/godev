package main

import (
	"fmt"
	"go_code/chatroom/client/process"
	"os"
)

//定义两个变量,一个是用户 id,一个用户 password
var (
	userId   int
	userPwd  string
	userName string
)

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	// var loop = true
	for {
		fmt.Println("------------------------欢迎登陆多人聊天系统-----------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择 1-3:")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的 id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			//完成登陆
			//创建一个 UserProcess
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户 id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户昵称:")
			fmt.Scanf("%s\n", &userName)
			//调用 UserProcess,完成注册请求
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入错误,请重新输入...")
		}

	}
}
