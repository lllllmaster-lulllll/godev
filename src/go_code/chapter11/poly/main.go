package main

import (
	"fmt"
)

//声明 定义一个接口
type Usb interface {
	//两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
	Name string
}

//让 Phone 实现 USB 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作...")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作...")
}

type Camera struct {
	Name string
}

//让 Camera 实现 Usb 接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

func main() {
	//定义一个 Usb 接口数组,可以存放 Phone 和 Camera的结构体变量
	var usbArr [3]Usb
	fmt.Println(usbArr)
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"huawei"}
	usbArr[2] = Camera{"canon"}
	fmt.Println(usbArr)
}
