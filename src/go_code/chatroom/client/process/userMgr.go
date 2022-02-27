package process

import (
	"fmt"
	"go_code/chatroom/client/model"
	"go_code/chatroom/common/message"
)

//客户端要维护的 map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //在用户登陆成功后,完成对 CurUser 初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历一把 onlineUsers
	fmt.Println("当前在线用户列表:")
	for id := range onlineUsers {
		fmt.Println("用户 id:\t", id)
	}
}

//编写方法,处理返回的NotifyUserStatusMes
func UpdateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
