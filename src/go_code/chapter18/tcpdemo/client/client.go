package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.69.169:9876")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Println("conn 成功", conn)
	//功能一:客户端发送单行数据,输入 exit 就退出
	for {
		reader := bufio.NewReader(os.Stdin) //标准输入
		//从终端读取一行用户输入,并发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
		}
		fmt.Printf("line=%v,len(line)=%v\n", line, len(line))
		newline := strings.Trim(line, "\r\n")
		if newline == "exit" {
			return
		}
		//再将 line 发送给服务器
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err:", err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据,并退出\n", n)
	}
}
