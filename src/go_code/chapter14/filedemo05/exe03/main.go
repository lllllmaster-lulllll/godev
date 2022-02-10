package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filepath := "/Users/tal/iceCode/godev/abc.txt"
	//打开已经存在的文件
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}
	defer file.Close()
	str := "BTBU 冰墩墩\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	//因为 writer 是带缓存的,因此在调用 WriterString 方法时,
	//内容是先写到缓存的,调用 Flush 方法,将缓存写入磁盘
	writer.Flush()
}
