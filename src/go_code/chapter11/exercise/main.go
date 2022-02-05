package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//1. 声明 Hero 结构体
type Hero struct {
	Name string
	Age  int
}

//声明一个 Hero 结构体切片类型
type HeroSlice []Hero

//3.实现 Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less 方法决定你使用什么标准进行排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	// 等价于下面这句
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//要求对 intSlice 切片进行排序
	//1.冒泡排序
	//2.也可以使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	//调用 sort.Sort
	sort.Sort(heroes)
	fmt.Println("-------------------排序后-----------------------")
	for _, v := range heroes {
		fmt.Println(v)
	}

}
