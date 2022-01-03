package main

import "fmt"

func main() {
	var a int=21
	var b int=10
	var c int
	c=a+b
	fmt.Printf("第1行 - c 的值为 %d\n",c)
	c=a-b
	fmt.Printf("第2行 - c 的值为 %d\n",c)
	c=a*b
	fmt.Printf("第3行 - c 的值为 %d\n",c)
	c=a/b
	fmt.Printf("第4行 - c 的值为 %d\n",c)
	c=a%b
	fmt.Printf("第5行 - c 的值为 %d\n",c)
	a++
	fmt.Printf("第6行 - c 的值为 %d\n",a)
	a=21
	a--
	fmt.Printf("第7行 - c 的值为 %d\n",a)
}

