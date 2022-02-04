package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (person Person) speak() {
	fmt.Println(person.Name, "是一个好人")
}
func (person Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println("1-1000的和为:", res)
}
func (person Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println("1-n 的和为:", res)
}
func (person Person) getSum(n1 int, n2 int) int {

	return n1 + n2
}

//给person类型绑定一个方法
func (person Person) test() {
	fmt.Println("test()", person.Name)
}
func main() {
	var p Person
	p.Name = "Tom"
	p.test()
	fmt.Println()
	p.speak()
	p.jisuan()
	p.jisuan2(89)
	res := p.getSum(2, 3)
	fmt.Println("getSum的结果为:", res)
}
