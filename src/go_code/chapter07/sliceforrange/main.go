package main

import (
	"fmt"
)

func main() {
	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4] //20.30.40
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}
	//for-range方式遍历
	for index, v := range slice {
		fmt.Printf("slice[%v]=%v\n", index, v)
	}
	//用append内置函数,可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	//通过append直接给slice3追加具体元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3=", slice3)
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3=", slice3)
	//切片的拷贝
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 []int = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)

}
