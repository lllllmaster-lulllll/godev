package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

//编写一个函数 serverProcessLogin 函数,专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//core code
	//1. 先从 mes 中取出mes.Data,并直接反序列化成 loginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json Unmarshal fail err=", err)
		return
	}
	//声明 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//声明 LoginResMes
	var loginResMes message.LoginResMes

	//如果用户的 id=100,密码是 123456 认为合法,否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//illegle
		loginResMes.Code = 500 //500 状态码,表示用户不存在
		loginResMes.Error = "用户不存在,请注册再使用..."
	}
	//将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("loginResMes marshal fail", err)
		return
	}
	//将 data 赋值给resMes
	resMes.Data = string(data)

	//对resMes 进行序列化,准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//发送 data, 将发送 data 的函数封装到 writePkg 函数
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
