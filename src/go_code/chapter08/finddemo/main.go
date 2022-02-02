package main

import "fmt"

func main() {
	//1) 有一个数列：白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王
	//猜数游戏：从键盘中任意输入一个名称，判断数列中是否包含此名称【顺序查找】
	//1 定义一个数组
	//2 从控制台接收一个名词,一次比较

	//code
	names := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}
	var heroName = ""
	fmt.Scanln(&heroName)
	//顺序查找的方式 -> 1
	// for i := 0; i < len(names); i++ {
	// 	if heroName == names[i] {
	// 		fmt.Printf("找到%v, 下标%v\n", heroName, i)
	// 	} else if i == (len(names) - 1) {
	// 		fmt.Printf("没有找到%v \n", heroName)
	// 	}
	// }
	//index进行记录
	index := -1
	//顺序查找的方式 -> 1
	for i := 0; i < len(names); i++ {
		if heroName == names[i] {
			index = i
		}
	}
	if index == -1 {
		fmt.Printf("没有找到%v \n", heroName)
	} else {
		fmt.Printf("找到%v, 下标%v\n", heroName, index)
	}

}
