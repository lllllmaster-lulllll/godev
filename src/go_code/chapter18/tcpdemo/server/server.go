package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端数据
	defer conn.Close() //关闭 conn
	fmt.Printf("服务器在等待客户端 %s 发送信息", conn.RemoteAddr().String())
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//conn.Read(buf)
		//等待客户端通过 conn 发送信息
		//如果客户端没有 write 发送,那么协程阻塞

		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("客户端退出")
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:9876")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close() //延时关闭 listen
	//循环等待客户端来链接我
	for {
		//等待客户端连接
		fmt.Println("等待客户端来连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Println("连接成功")
			fmt.Printf("Accept() suc conn=%v 客户端 ip=%v\n", conn, conn.RemoteAddr())
		}
		//这里准备起一个协程,为客户端服务
		go process(conn)
	}

	// fmt.Printf("listen suc=%v\n", listen)

}
