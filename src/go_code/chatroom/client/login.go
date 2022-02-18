package main

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
)

//写一个函数,完成登陆
func login(userId int, userPwd string) (err error) {
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
	err = utils.WritePkg(conn, data)
	if err != nil {
		fmt.Println("WritePkg fail")
		return
	}
	//处理服务器返回的数据

	mes, err = utils.ReadPkg(conn)
	if err != nil {
		fmt.Println("ReadPkg fail", err)
		return
	}
	//将 mes 的 Data 部分反序列化成 LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
