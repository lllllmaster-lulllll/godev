package main

import "fmt"

//func main(){
//	sum:=0
//	for i:=0;i<=10;i++{
//		sum+=i
//	}
//	fmt.Println(sum)
//}
//func main(){
//	sum:=1
//	for ;sum<=10;{
//		sum+=sum
//	}
//	fmt.Println(sum)
//	sum=1
//	for sum<=10{
//		sum+=sum
//	}
//	fmt.Println(sum)
//}
//func main(){
//	strings := []string{"google","runoob"}
//	for i, s := range strings{
//		fmt.Println(i, s)
//	}
//	numbers:=[6]int{1,2,3,5}
//	for i,x:=range numbers{
//		fmt.Printf("第%d位 x 的值 = %d\n",i,x)
//	}
//}
//func main() {
//
//	// 不使用标记
//	fmt.Println("---- break ----")
//	for i := 1; i <= 3; i++ {
//		fmt.Printf("i: %d\n", i)
//		for i2 := 11; i2 <= 13; i2++ {
//			fmt.Printf("i2: %d\n", i2)
//			break
//		}
//	}
//
//	// 使用标记
//	fmt.Println("---- break label ----")
//mmmmm:for i := 1; i <= 3; i++ {
//		fmt.Printf("i: %d\n", i)
//		for i2 := 11; i2 <= 13; i2++ {
//			fmt.Printf("i2: %d\n", i2)
//			break mmmmm
//		}
//	}
//}
func main(){
	var a int=10
	LOOP:for a<20{
		if a==15{
			a=a+1
			goto LOOP
		}
		fmt.Printf("a的值为：%d\n",a)
		a++
	}
}
