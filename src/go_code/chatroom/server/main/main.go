package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//读客户端发送的消息
	defer conn.Close()

	//循环的读取客户端发送的信息
	//这里调用总控,创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程出错 err=", err)
		return
	}
}
func main() {
	//提示信息
	fmt.Println("服务器[面向对象版本]在 8889 端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	//一旦监听成功,就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen,Accept err=", err)
			return
		}
		//一旦连接成功,则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
