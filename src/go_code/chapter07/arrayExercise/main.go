package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//创建一个byte类型的26个元素的数组，分别放置‘A’-‘Z’。
	//使用for循环访问所有元素并打印出来。提示：字符数组运算‘A’+1->B
	//思路
	//1. 声明一个数组 var myChars [26]byte
	//2. 使用for循环,利用字符可以运算的特点来赋值
	//code
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)

	}
	for j, v := range myChars {
		fmt.Printf("第%d个元素字符为%c\n", j, v)
	}
	//请求出一个数组的最大值，并得到对应的下标
	//思路1. 声明一个数组 var intArr [5]int=[...]int{1,-1,9,90,11}
	var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	var maxValue = intArr[0]
	var indexOfMaxValue = 0
	for index, value := range intArr {
		if value > maxValue {
			maxValue = value
			indexOfMaxValue = index
		}
	}
	fmt.Printf("最大值为%d,其对应的下标为%d\n", maxValue, indexOfMaxValue)
	//请求出一个数组的和和平均值。for-range
	var intArr2 [5]int = [...]int{1, -1, 9, 90, 12}
	sum := 0.0
	for _, value := range intArr2 {
		sum += float64(value)
	}
	fmt.Printf("数组的平均值为%v\n", sum/float64(len(intArr2)))

	//要求:随机生成五个数,并将其反转打印
	//思路
	//1. 随机生成五个数,rand.Intn()函数
	//2. 得到随机数后,就放到一个数组中 int数组
	//3. 反转打印,交换的次数是 len/2
	var intArr3 [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100) //0<=n<=100
	}
	fmt.Println(intArr3)
	var temp int = 0
	for i := 0; i < len(intArr3)/2; i++ {
		temp = intArr3[i]
		intArr3[i] = intArr3[len(intArr3)-1-i]
		intArr3[len(intArr3)-1-i] = temp
	}
	fmt.Println(intArr3)

}
