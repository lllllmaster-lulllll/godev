package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	//打开已经存在的文件:将 abc.txt 导入到 kkk.txt
	file1path := "/Users/aurora/iceCode/godev/abc.txt"
	file2path := "/Users/aurora/iceCode/godev/kkk.txt"
	data, err := ioutil.ReadFile(file1path)
	if err != nil {
		//说明读取文件有错误
		fmt.Printf("read file err:%v\n", err)
	}
	err = ioutil.WriteFile(file2path, data, 0666)
	if err != nil {
		fmt.Printf("write file err:%v\n", err)
	}
	fmt.Println(PathExists(file1path))

}
