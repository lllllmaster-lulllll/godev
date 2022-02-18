package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过 go 向 redis 写入数据和读取数据
	//1. 链接到 redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer conn.Close()
	fmt.Println("conn succ...", conn)
	_, err = conn.Do("Set", "name", "tomjerry猫猫")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//通过go 从 redis 读取数据
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
	}
	//返回的 r 是 interface{},需要转换一下
	fmt.Println("操作 ok", r)
}
