package main

import "fmt"

// func test01(arr [3]int) {
// arr[0] = 88
// }
func test02(arr *[3]int) {
	(*arr)[0] = 88
	// arr[0] = 88
}
func main() {
	var arr [3]int = [3]int{11, 22, 33}
	test02(&arr)
	fmt.Println("main  arr=", arr)

}
