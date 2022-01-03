package main

import "fmt"

//func main(){
//	var a int=100
//	var b int=200
//	fmt.Printf("交换前，a 的值：%d\n",a)
//	fmt.Printf("交换前，b 的值：%d\n",b)
//	swap(&a, &b)
//	fmt.Printf("交换后，a 的值：%d\n",a)
//	fmt.Printf("交换后，b 的值：%d\n",b)
//}
//
//func swap(x *int,y *int){
//	var temp int
//	temp=*x
//	*x=*y
//	*y=temp
//}

//func main(){
//	getSquareRoot:=func(x float64) float64{
//		return math.Sqrt(x)
//	}
//	fmt.Println(getSquareRoot(9))
//}
//type cb func(int) int
//func main(){
//	testCallBack(1,callBack)
//	testCallBack(2, func(x int) int {
//		fmt.Printf("我是回调，x:%d\n",x)
//		return x
//	})
//}
//func testCallBack(x int,f cb){
//	f(x)
//}
//func callBack(x int) int{
//	fmt.Printf("我是回调，x：%d\n",x)
//	return x
//}
//func getSequence() func() int{
//	i:=0
//	return func() int {
//		i+=1
//		return i
//	}
//}
//func main(){
//	nextNumber:=getSequence()
//	fmt.Println(nextNumber())
//	nextNumber2:=nextNumber()
//	fmt.Println(nextNumber2)
//	nextNumber2=nextNumber()
//	fmt.Println(nextNumber2)
//	nextNumber1:=getSequence()
//	fmt.Println(nextNumber1())
//	fmt.Println(nextNumber1())
//
//}

type Circle struct{
	radius float64
}
func main(){
	var c1 Circle
	c1.radius=10
	fmt.Println("圆的面积 = ", c1.getArea())
}
func (c Circle) getArea() float64{
	return 3.14*c.radius*c.radius
}
