package model

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"

	"github.com/garyburd/redigo/redis"
)

//服务器启动后,初始化一个 userDao 实例
//做成全局变量,在需要和 redis 交互时,直接使用
var (
	MyUserDao *UserDao
)

//定义一个 UserDao 结构体
//完成对 User结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式,创建一个 UserDao 实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//思考需要哪些方法
//1. 根据用户 id 返回一个 User 实例+err
func (ptr *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定 id 去 redis 查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示在 users 哈希中
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//这里需要把 res 反序列化乘 User 实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//完成登陆的校验
//Login完成对用户的验证
//如果用户的 id 和 pwd 都正确,则返回一个 user 实例
//如果用户的 id 或 pwd 有错误,则返回对应的错误信息
func (ptr *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 先从 UserDao 的连接池中取出一根连接
	conn := ptr.pool.Get()
	defer conn.Close()
	user, err = ptr.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
func (ptr *UserDao) Register(user *message.User) (err error) {
	// 先从 UserDao 的连接池中取出一根连接
	conn := ptr.pool.Get()
	defer conn.Close()
	_, err = ptr.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这时,说明 id 在 redis 还没有,则可以完成注册
	data, err := json.Marshal(user) //序列化
	if err != nil {
		return
	}
	//入库
	_, err = conn.Do("Hset", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}
