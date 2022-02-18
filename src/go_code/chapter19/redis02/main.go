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
	_, err = conn.Do("HSet", "user01", "name", "康楠")
	if err != nil {
		fmt.Println("Hset err=", err)
		return
	}
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}
	//通过go 从 redis 读取数据
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("get err=", err)
	}
	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("get err=", err)
	}
	//返回的 r 是 interface{},需要转换一下
	fmt.Printf("r1=%v r2=%v\n", r1, r2)
}
