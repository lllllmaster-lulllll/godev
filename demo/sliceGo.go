package main

import "fmt"

func main() {
	//var numbers=make([]int,3,5)
	//printSlice(numbers)
	//var numbers []int
	//printSlice(numbers)
	//if(numbers==nil){
	//	fmt.Printf("切片是空的")
	//}
	//numbers:=[]int{0,1,2,3,4,5,6,7,8}
	//printSlice(numbers)
	//fmt.Println("numbers==",numbers)
	//
	//fmt.Println("numbers[1:4]==",numbers[1:4])
	//fmt.Println("numbers[:3]==",numbers[:3])
	//fmt.Println("numbers[4:]==",numbers[4:])
	//
	//numbers1:=make([]int,0,5)
	//printSlice(numbers1)
	//numbers2:=numbers[:2]
	//printSlice(numbers2)
	//numbers3:=numbers[2:5]
	//printSlice(numbers3)
	var numbers []int
	printSlice(numbers)
	numbers=append(numbers,0)
	printSlice(numbers)

	numbers=append(numbers,1)
	printSlice(numbers)
	numbers=append(numbers,2,3,4)
	printSlice(numbers)
	numbers1:=make([]int,len(numbers),(cap(numbers))*2)
	copy(numbers1,numbers)
	printSlice(numbers1)

}
func printSlice(x [] int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
