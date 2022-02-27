package main

import (
	"fmt"

	"go_code/chatroom/common/message"
	process2 "go_code/chatroom/server/process"
	"go_code/chatroom/server/utils"
	"io"
	"net"
)

//先创建一个 Processor 的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个serverProcessMes函数
//功能:根据客户端发送消息的种类不同,决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//看看是否接收到客户端发送的群发消息
	fmt.Println("mes=", mes)
	switch mes.Type {
	case message.LoginMesType:
		//处理登陆的逻辑
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)

	case message.RegisterMesType:
		//处理注册逻辑
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//创建一个 SmsProcess 实例完成转发群聊消息
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在,无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {
	for {
		//读取数据包,并进行封装
		//创建 Transfer 实例完成读包
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出,服务器端也正常退出")
				return err
			} else {
				fmt.Println("readPkg err=", err)
			}

		}
		// fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}

}
