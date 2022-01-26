package main

import (
	"fmt"
)

func main() {
	//for-range遍历数组
	var heros [3]string = [3]string{"宋江", "无用", "武松"}
	for i, v := range heros {
		fmt.Printf("i = %v, v=%v\n", i, v)

	}
	for _, v := range heros {
		fmt.Printf("i = %v, v=%v\n", i, v)

	}
}
