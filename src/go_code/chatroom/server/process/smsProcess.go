package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
}

//写方法转发消息{
func (ptr *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的 onlineUsers map[int]*UserProcess,
	//将消息转发取出

	//取出 mes 的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		//这里,还需要过滤掉自己,不要再发给自己
		if id == smsMes.UserId {
			continue
		}
		ptr.SendMesToEachOnlineUser(data, up.Conn)
	}
}
func (ptr *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个 Transfer 实例,发送 data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("妆发消息失败 err=", err)
	}
}
