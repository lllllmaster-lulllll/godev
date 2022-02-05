package main

type A struct {
	Name string
	Age  int
}
type B struct {
	Name string
	Age  int
}
type C struct {
	A
	B
}
type D struct {
	a A //嵌套有名结构体
}

func main() {
	//如果 D中是一个有名结构体,则访问有名结构体的字段时,
	//就必须带上有名结构体的名字,比如 d.a.Name
	var d D
	d.a.Name = "jack"
}
