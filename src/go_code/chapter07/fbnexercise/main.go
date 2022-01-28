package main

import "fmt"

func fbn(n int) []uint64 {
	fbnSlice := make([]uint64, n)
	//第一个数和第二个数的斐波那契为1
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	//for循环来存放斐波那契数列
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-2] + fbnSlice[i-1]
	}
	return fbnSlice
}
func main() {
	// 说明：编写一个函数 fbn(n int) ，要求完成
	// 1) 可以接收一个 n int .
	// 2）能够将斐波那契的数列放到切片中提示，
	// 斐波那契的数列形式 an[0] = 1; arr[1] = 1; an[2]=2; an[3] = 3; an[4]=5; an[5]=8
	fbnSlice := fbn(10)
	fmt.Println("fbnSlice=", fbnSlice)

}
