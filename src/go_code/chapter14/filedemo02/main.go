package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("/Users/tal/iceCode/godev/data.txt")
	if err != nil {
		fmt.Println("open file err:", err)
	}

	//关闭文件 当函数退出时,
	defer file.Close() //要及时的关闭 file
	//创建一个*Reader,是带缓冲的
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)

	}
	fmt.Println("文件读取结束...")
}
