package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
)

func outputGroupMes(mes *message.Message) { //这个地方的 mes 一定是 smsMes
	//显示即可
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		println("smsMes json.Unmarshal err", err)
		return
	}
	//显示信息
	info := fmt.Sprintf("用户 id:\t%d 对大家说:\t%s\n", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}
