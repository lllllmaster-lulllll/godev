package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用 ioutil.ReadFile 一次性将文件读取到位
	file := "/Users/tal/iceCode/godev/data.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err:%v\n", err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("content:%v\n", string(content))
	//没有显示 openfile 所以不需要显示的 close 文件
	//因为文件的打开和关闭被封装在函数中

}
