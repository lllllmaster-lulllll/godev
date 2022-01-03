package main

import "fmt"

func main() {
	//var countryCapitalMap map[string]string
	//countryCapitalMap=make(map[string]string)
	//countryCapitalMap["France"]="巴黎"
	//countryCapitalMap["Italy"]="罗马"
	//countryCapitalMap["Japan"]="东京"
	//countryCapitalMap["India"]="新德里"
	//
	//for country:=range countryCapitalMap{
	//	fmt.Println(country,"首都是：",countryCapitalMap[country])
	//}
	//capital,ok:=countryCapitalMap["America"]
	//if(ok){
	//	fmt.Println("America 的首都是：",capital)
	//}else{
	//	fmt.Println("America 的首都不存在")

	countryCapitalMap:=map[string]string{"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New delhi"}
	fmt.Println("原始地图")
	for country:=range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}
	delete(countryCapitalMap,"France")
	fmt.Println("法国条目被删除")
	fmt.Println("删除元素后地图")
	for country:=range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}
}

