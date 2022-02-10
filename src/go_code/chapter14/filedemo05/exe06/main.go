package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//编写一个函数,接收两个文件路径 srcFilePath dstFilePath
func CopyFile(dsfFilePath string, srcFilePath string) (written int64, err error) {
	srcfile, err := os.Open(srcFilePath)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer srcfile.Close()
	//通过 srcfile,获取到 Reader
	reader := bufio.NewReader(srcfile)

	//打开 dstfile
	dstfile, err := os.OpenFile(dsfFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	//通过 dstfile,获取到 writer
	writer := bufio.NewWriter(dstfile)
	return io.Copy(writer, reader)

}

func main() {
	//拷贝图片
	//调用 CopyFile 完成文件拷贝
	srcFile := "/Users/aurora/iceCode/godev/src/go_code/chapter14/gogo.jpg"
	dstFile := "/Users/aurora/iceCode/godev/src/go_code/chapter14/filedemo05/exe06/gogo.jpg"
	returecode, err := CopyFile(dstFile, srcFile)
	fmt.Println(returecode)
	if err != nil {
		fmt.Println("copy file err:", err)
	} else {
		fmt.Println("copy succeed")
	}

}
