package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

type UserProcess struct {
	//
}

func (ptr *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//连接到服务器
	conn, err := net.Dial("tcp", "localhost:8899")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//准备通过发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3. 创建一个 RegisterMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//序列化
	//4. 将 RegisterMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}
	//将 data 赋给 mes.Data 字段
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg fail")
		return
	}
	//处理服务器返回的数据

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg fail", err)
		return
	}
	//将 mes 的 Data 部分反序列化成 RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功,请重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}

//写一个函数,完成登陆
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//定协议
	// fmt.Printf(" userId = %d userPwd=%s \n", userId, userPwd)
	// return nil
	//连接到服务器
	conn, err := net.Dial("tcp", "localhost:8899")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//准备通过发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//3. 创建一个 LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. 将 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}
	//将 data 赋给 mes.Data 字段
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err=", err)
		return
	}
	//发送数据,
	//需要先获取 data 字符串长度,发送给服务器
	//先获取到 data 的长度->转成一个表示长度的切片
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg fail")
		return
	}
	//处理服务器返回的数据

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("ReadPkg fail", err)
		return
	}
	//将 mes 的 Data 部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化 CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		//显示当前在线用户列表,遍历 loginResMes.UsersId
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户 id:\t", v)
			//完成客户端的 onlineUsers 完成初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user

		}
		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯,如果服务器有数据推送给我们
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)
		//1. 显示我们的登录成功的菜单(循环)
		ShowMenu()
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
