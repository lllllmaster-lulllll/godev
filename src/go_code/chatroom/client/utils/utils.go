package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

type Transfer struct {
	//分析应该有的字段
	Conn net.Conn
	Buf  [8096]byte //传输时,使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("读取客户端发送的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	fmt.Println("读到的 buf=", this.Buf[:4])
	//根据 buf[0:4]转成一个 uint32 类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	//根据 pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	//把 pkgLen 反序列化成 message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		// err = errors.New("Unmarshal error")
		return
	}
	return mes, nil
}

//发送 data 的函数进行封装 writePkg
func (this *Transfer) WritePkg(data []byte) (err error) {
	//发送数据,
	//需要先获取 mesData 字符串长度,发送给服务器
	//先获取到 mesData 的长度->转成一个表示长度的切片
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("长度 conn.Write(bytes) fail", err)
		return
	}
	//发送消息本身
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("消息本身 conn.Write(data) fail", err)
		return
	}
	return
}
