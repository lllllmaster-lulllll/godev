package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("/Users/tal/iceCode/godev/data.txt")
	if err != nil {
		fmt.Println("open file err:", err)
	}
	//输出文件
	fmt.Printf("file=%v\n", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("open file err:", err)
	}
}
