package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := "/Users/tal/iceCode/godev/abc.txt"
	//打开已经存在的文件
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}
	defer file.Close()
	//先读取原来文件的内容,并显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	str := "hello 北京\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	//因为 writer 是带缓存的,因此在调用 WriterString 方法时,
	//内容是先写到缓存的,调用 Flush 方法,将缓存写入磁盘
	writer.Flush()
}
