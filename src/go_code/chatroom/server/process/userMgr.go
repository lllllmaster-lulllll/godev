package process2

import "fmt"

//UserMgr 实例在服务器端有且只有一个
//在很多地方都会使用到
//将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对 userMgr 初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对 onlineUsers 添加
func (ptr *UserMgr) AddOnlineUser(up *UserProcess) {
	ptr.onlineUsers[up.UserId] = up
}

//删除
func (ptr *UserMgr) DelOnlineUser(userId int) {
	delete(ptr.onlineUsers, userId)
}

//返回当前所有在线的用户
func (ptr *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return ptr.onlineUsers
}

//根据 id 返回对应的值
func (ptr *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//如何从 map 取出一个值,带检测方式
	up, ok := ptr.onlineUsers[userId]
	if !ok { //说明,你要查找的这个用户,当前不在线
		err = fmt.Errorf("用户 %d 不存在", userId)
		return
	}
	return
}
