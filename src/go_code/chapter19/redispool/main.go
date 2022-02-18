package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//定义一个全局 pool
var pool *redis.Pool

//当启动程序时,就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	//先从 pool 取出一个链接
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "name", "汤姆猫~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println("获取到的值为:", r)
	pool.Close()
	conn2 := pool.Get()
	_, err = conn2.Do("set", "name2", "汤姆猫2~~")
	if err != nil {
		fmt.Println("conn2.Do err=", err)
		return
	}
	r2, err := redis.String(conn2.Do("get", "name2"))
	if err != nil {
		fmt.Println("conn2.Do err=", err)
		return
	}
	fmt.Println("获取到的值为:", r2)

}
