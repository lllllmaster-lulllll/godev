package main

import (
	"fmt"
)

func main() {
	var intArr [3]int
	//定义完了以后,数组各个元素有默认值0
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p\n",
		&intArr, &intArr[0], &intArr[1], &intArr[2])
}
