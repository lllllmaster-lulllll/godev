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
	//从终端输入五个成绩,保存到float64的数组里面,然后输出
	var scores [5]float64
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请输入第%d个元素的值\n", i+1)
		fmt.Scanln(&scores[i])
	}
	for i := 0; i < len(scores); i++ {
		fmt.Printf("score[%d]的值为%v\n", i, scores[i])

	}
	var myVal float64 = 9
	fmt.Printf("myVal 地址 %p, 地址2 %x\n", &myVal, &myVal)
	//四种初始化数组的方式
	var numsArray01 [3]int = [3]int{1, 2, 3}
	var numsArray02 = [3]int{1, 2, 3}
	var numsArray03 = [...]int{6, 7, 8}
	//可以指定元素值对应的下标
	var names = [3]string{1: "tom", 2: "marie", 0: "jack"}
	var count = [...]int{1: 11111, 0: 222222, 2: 33333333}
	fmt.Println(numsArray01)
	fmt.Println(numsArray02)
	fmt.Println(numsArray03)
	fmt.Println(names)
	fmt.Println(count)
}
