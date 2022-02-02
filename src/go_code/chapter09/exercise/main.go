package main

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string, name string) {
	//判断是否有名字为name的用户
	if users[name] != nil {
		//有这个用户
		users[name]["pwd"] = "888888"
	} else {
		//没有
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "昵称~" + name
		users[name]["nickname"] = "昵称~" + name
	}
}

func main() {
	// 1）使用map[string]map[string]sting的map类型
	// 2）key:表示用户名，是唯一的，不可以重复
	// 3）如果某个用户名存在，就将其密码修改"888888"， 如果不存在就增加这个用户 信息,（包括昵称nickname和密码pwd）
	// 4）编写一个函数 modifyUser(users map[string]map[string]sting,name string)完 成上述功能
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "000000"
	users["smith"]["nickname"] = "昵称~" + "smith"
	modifyUser(users, "Tom")
	modifyUser(users, "Mary")
	modifyUser(users, "smith")
	fmt.Println(users)
}
