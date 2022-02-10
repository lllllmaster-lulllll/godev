package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义一个结构体,用于保存统计结构
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	//思路:打开一个文件,创建一个 Reader
	//每读取一行,就去统计该行有多少个 英文,数字,空格和其他字符
	//然后将结果保存到一个结构体
	filepath := "/Users/aurora/iceCode/godev/data.txt"
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}
	defer file.Close()

	//定义一个 CharCount 实例
	var count CharCount
	//创建一个 Reader
	reader := bufio.NewReader(file)

	//开始循环读取 file 的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历 str,进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透处理
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}

	}
	//输出统计结果
	fmt.Printf("字符的个数为:%v 数字的个数为:%v 空格的个数为:%v 其他字符个数为:%v\n",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}
