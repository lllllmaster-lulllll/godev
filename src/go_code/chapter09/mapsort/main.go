package main

import (
	"fmt"
	"sort"
)

func main() {
	//map排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	for k, v := range map1 {
		fmt.Printf("打印k=%v,v=%v\n", k, v)
	}
	//如果按照map的key的顺序进行排序输出
	//1. 先将map的key放入到一个切片中,对切片进行排序
	//2. 对切片排序
	//3. 遍历切片,然后按照key来输出map的值
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Printf("map[%v]=%v\n", v, map1[v])
	}
}
